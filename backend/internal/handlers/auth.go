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
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	BloodGroup string `json:"bloodGroup" binding:"required"`
	Gender     string `json:"gender"`   // Added Gender
	Birthday   string `json:"birthday"` // Changed to string for manual parsing
	Division   string `json:"division"`
	District   string `json:"district"`
	Upazila    string `json:"upazila"`
	Area       string `json:"area"`
	City       string `json:"city"`
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Generate UniqueID: BD-XXXXXX
	userID := utils.GenerateUniqueID()
	verificationToken := utils.GenerateUniqueID() // Reusing the unique ID gen for token simplicity, or can use better entropy

	user := models.User{
		ID:                userID,
		Email:             input.Email,
		PasswordHash:      string(hashedPassword),
		Role:              models.RoleDonor,
		VerificationToken: verificationToken,
		IsVerified:        false,
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
	}

	// Check if phone number already exists
	var existingProfile models.DonorProfile
	if err := config.DB.Where("phone = ?", input.Phone).First(&existingProfile).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number is already registered"})
		return
	}

	// Transaction to create User and DonorProfile
	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user. Please try another email or login."})
		return
	}

	// Parse birthday if provided
	var birthdayPtr *time.Time
	if input.Birthday != "" {
		t, err := time.Parse("2006-01-02", input.Birthday)
		if err == nil {
			birthdayPtr = &t
		}
	}

	profile := models.DonorProfile{
		UserID:      userID,
		Name:        input.Name,
		Phone:       input.Phone,
		BloodGroup:  input.BloodGroup,
		Gender:      input.Gender,   // Save Gender
		Birthday:    birthdayPtr,    // Save parsed Birthday
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
	tx.Commit()

	// Send verification email
	utils.SendVerificationEmail(user.Email, verificationToken)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully. Please check your email for verification link."})
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

	if !user.IsVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Please verify your email before logging in"})
		return
	}

	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	var user models.User
	if err := config.DB.Where("verification_token = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid or expired token"})
		return
	}

	user.IsVerified = true
	user.VerificationToken = ""
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Don't reveal if user exists for security, just send same response
		c.JSON(http.StatusOK, gin.H{"message": "If this email is registered, a reset link has been sent"})
		return
	}

	resetToken := utils.GenerateUniqueID()
	expires := time.Now().Add(1 * time.Hour)

	user.ResetToken = resetToken
	user.ResetTokenExpires = &expires
	config.DB.Save(&user)

	utils.SendPasswordResetEmail(user.Email, resetToken)

	c.JSON(http.StatusOK, gin.H{"message": "If this email is registered, a reset link has been sent"})
}

func ResetPassword(c *gin.Context) {
	var input struct {
		Token    string `json:"token" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("reset_token = ? AND reset_token_expires > ?", input.Token, time.Now()).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(hashed)
	user.ResetToken = ""
	user.ResetTokenExpires = nil
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
