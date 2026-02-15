package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
			userMap["last_donation"] = user.DonorProfile.LastDonationDate
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
