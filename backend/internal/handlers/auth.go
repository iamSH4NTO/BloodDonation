package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"blood-donor-system/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	BloodGroup string `json:"bloodGroup" binding:"required"`
	Division   string `json:"division"`
	District   string `json:"district"`
	Upazila    string `json:"upazila"`
	Area       string `json:"area"`
	PostalCode string `json:"postalCode"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         models.RoleDonor,
	}

	// Transaction to create User and DonorProfile
	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	profile := models.DonorProfile{
		UserID:      user.ID,
		Name:        input.Name,
		Phone:       input.Phone,
		BloodGroup:  input.BloodGroup,
		District:    input.District,
		Upazila:     input.Upazila,
		City:        input.Upazila, // Mapping Upazila to City as they are often interchangeable in this context
		AreaVillage: input.Area,
		PostalCode:  input.PostalCode,
	}

	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, string(user.Role))

	c.JSON(http.StatusOK, gin.H{"token": token})
}
