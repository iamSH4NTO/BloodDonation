package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetDashboardStats(c *gin.Context) {
	var totalDonors int64
	var totalDonations int64
	var totalUsers int64

	config.DB.Model(&models.DonorProfile{}).Count(&totalDonors)
	config.DB.Model(&models.Donation{}).Count(&totalDonations)
	config.DB.Model(&models.User{}).Count(&totalUsers)

	// Calculate recent donations (last 30 days) - optional enhancement
	// var recentDonations int64
	// lastMonth := time.Now().AddDate(0, -1, 0)
	// config.DB.Model(&models.Donation{}).Where("date >= ?", lastMonth).Count(&recentDonations)

	c.JSON(http.StatusOK, gin.H{
		"total_donors":    totalDonors,
		"total_donations": totalDonations,
		"total_users":     totalUsers,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	// Preload profile to get name/blood group
	if err := config.DB.Preload("DonorProfile").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Transform for frontend table
	var userList []gin.H
	for _, user := range users {
		userMap := gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
			"is_active":  user.IsActive,
		}
		if user.DonorProfile != nil {
			userMap["name"] = user.DonorProfile.Name
			userMap["blood_group"] = user.DonorProfile.BloodGroup
			userMap["phone"] = user.DonorProfile.Phone
			userMap["district"] = user.DonorProfile.District
			userMap["last_donation_date"] = user.DonorProfile.LastDonationDate
			userMap["gender"] = user.DonorProfile.Gender
			userMap["birthday"] = user.DonorProfile.Birthday
			userMap["upazila"] = user.DonorProfile.Upazila
			userMap["city"] = user.DonorProfile.City
			userMap["area_village"] = user.DonorProfile.AreaVillage
			userMap["postal_code"] = user.DonorProfile.PostalCode
			userMap["google_map_link"] = user.DonorProfile.GoogleMapLink
			userMap["is_available"] = user.DonorProfile.IsAvailable
		} else {
			userMap["name"] = "N/A"
			userMap["blood_group"] = "N/A"
		}
		userList = append(userList, userMap)
	}

	c.JSON(http.StatusOK, userList)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

type CreateUserInput struct {
	Email            string     `json:"email" binding:"required,email"`
	Password         string     `json:"password" binding:"required,min=6"`
	Name             string     `json:"name" binding:"required"`
	Role             string     `json:"role" binding:"required"` // 'donor' or 'admin'
	Phone            string     `json:"phone" binding:"required"`
	BloodGroup       string     `json:"bloodGroup" binding:"required"`
	Gender           string     `json:"gender"`
	Birthday         *time.Time `json:"birthday"`
	District         string     `json:"district"`
	Upazila          string     `json:"upazila"`
	City             string     `json:"city"`
	AreaVillage      string     `json:"area_village"`
	PostalCode       string     `json:"postal_code"`
	GoogleMapLink    string     `json:"google_map_link"`
	LastDonationDate *time.Time `json:"last_donation_date"`
	IsAvailable      bool       `json:"is_available"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Validate Role
	role := models.RoleDonor
	if input.Role == "admin" {
		role = models.RoleAdmin
	}

	user := models.User{
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         role,
		IsActive:     true,
	}

	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user (Email might be taken)"})
		return
	}

	profile := models.DonorProfile{
		UserID:           user.ID,
		Name:             input.Name,
		Phone:            input.Phone,
		BloodGroup:       input.BloodGroup,
		Gender:           input.Gender,
		Birthday:         input.Birthday,
		District:         input.District,
		Upazila:          input.Upazila,
		City:             input.City,
		AreaVillage:      input.AreaVillage,
		PostalCode:       input.PostalCode,
		GoogleMapLink:    input.GoogleMapLink,
		LastDonationDate: input.LastDonationDate,
		IsAvailable:      input.IsAvailable,
	}

	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

type UpdateUserInput struct {
	Name             string     `json:"name"`
	Role             string     `json:"role"`
	IsActive         *bool      `json:"is_active"` // Use pointer to handle false
	Phone            string     `json:"phone"`
	BloodGroup       string     `json:"bloodGroup"`
	Gender           string     `json:"gender"`
	Birthday         *time.Time `json:"birthday"`
	District         string     `json:"district"`
	Upazila          string     `json:"upazila"`
	City             string     `json:"city"`
	AreaVillage      string     `json:"area_village"`
	PostalCode       string     `json:"postal_code"`
	GoogleMapLink    string     `json:"google_map_link"`
	LastDonationDate *time.Time `json:"last_donation_date"`
	IsAvailable      *bool      `json:"is_available"` // Pointer for false
	Password         string     `json:"password"`     // Optional
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update User fields
	if input.Role != "" {
		if input.Role == "admin" {
			user.Role = models.RoleAdmin
		} else {
			user.Role = models.RoleDonor
		}
	}
	if input.IsActive != nil {
		user.IsActive = *input.IsActive
	}
	if input.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		user.PasswordHash = string(hashed)
	}
	config.DB.Save(&user)

	// Update Profile fields
	var profile models.DonorProfile
	if err := config.DB.Where("user_id = ?", user.ID).First(&profile).Error; err == nil {
		if input.Name != "" {
			profile.Name = input.Name
		}
		if input.Phone != "" {
			profile.Phone = input.Phone
		}
		if input.BloodGroup != "" {
			profile.BloodGroup = input.BloodGroup
		}
		if input.Gender != "" {
			profile.Gender = input.Gender
		}
		if input.Birthday != nil {
			profile.Birthday = input.Birthday
		}
		if input.District != "" {
			profile.District = input.District
		}
		if input.Upazila != "" {
			profile.Upazila = input.Upazila
		}
		if input.City != "" {
			profile.City = input.City
		}
		if input.AreaVillage != "" {
			profile.AreaVillage = input.AreaVillage
		}
		if input.PostalCode != "" {
			profile.PostalCode = input.PostalCode
		}
		if input.GoogleMapLink != "" {
			profile.GoogleMapLink = input.GoogleMapLink
		}
		if input.LastDonationDate != nil {
			profile.LastDonationDate = input.LastDonationDate
		}
		if input.IsAvailable != nil {
			profile.IsAvailable = *input.IsAvailable
		}
		config.DB.Save(&profile)
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
