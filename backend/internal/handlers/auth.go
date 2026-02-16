package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"blood-donor-system/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email      string     `json:"email" binding:"required,email"`
	Password   string     `json:"password" binding:"required,min=6"`
	Name       string     `json:"name" binding:"required"`
	Phone      string     `json:"phone" binding:"required"`
	BloodGroup string     `json:"bloodGroup" binding:"required"`
	Gender     string     `json:"gender"`   // Added Gender
	Birthday   *time.Time `json:"birthday"` // Added Birthday
	Division   string     `json:"division"`
	District   string     `json:"district"`
	Upazila    string     `json:"upazila"`
	Area       string     `json:"area"`
	City       string     `json:"city"`
	PostalCode string     `json:"postalCode"`
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

	// Generate UniqueID: BD-XXXXXX
	userID := utils.GenerateUniqueID()

	user := models.User{
		ID:           userID,
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
		UserID:      userID,
		Name:        input.Name,
		Phone:       input.Phone,
		BloodGroup:  input.BloodGroup,
		Gender:      input.Gender,   // Save Gender
		Birthday:    input.Birthday, // Save Birthday
		Division:    input.Division, // Save Division
		District:    input.District,
		Upazila:     input.Upazila,
		City:        input.City,
		AreaVillage: input.Area,
		PostalCode:  input.PostalCode,
	}

	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	tx.Commit()

	// Register locations in Registry
	if profile.AreaVillage != "" {
		var existing models.LocationRegistry
		if err := config.DB.Where("name = ? AND type = ? AND district = ?", profile.AreaVillage, "village", profile.District).First(&existing).Error; err != nil {
			config.DB.Create(&models.LocationRegistry{
				Name:     profile.AreaVillage,
				Type:     "village",
				District: profile.District,
				Upazila:  profile.Upazila,
			})
		}
	}
	if profile.City != "" {
		var existing models.LocationRegistry
		if err := config.DB.Where("name = ? AND type = ? AND district = ?", profile.City, "city", profile.District).First(&existing).Error; err != nil {
			config.DB.Create(&models.LocationRegistry{
				Name:     profile.City,
				Type:     "city",
				District: profile.District,
				Upazila:  profile.Upazila,
			})
		}
	}

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
