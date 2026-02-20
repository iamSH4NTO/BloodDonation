package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"blood-donor-system/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
	// Pagination and Search parameters
	page := 1
	limit := 20
	q := c.Query("q")

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}
	offset := (page - 1) * limit

	var users []models.User
	query := config.DB.Model(&models.User{}).Preload("DonorProfile")

	// Apply search filter if provided
	if q != "" {
		searchQuery := "%" + q + "%"
		// Search in User (email, ID) and DonorProfile (name, phone, location fields)
		query = query.Joins("LEFT JOIN donor_profiles ON donor_profiles.user_id = users.id").
			Where("users.email LIKE ? OR users.id LIKE ? OR donor_profiles.name LIKE ? OR donor_profiles.phone LIKE ? OR donor_profiles.district LIKE ? OR donor_profiles.city LIKE ? OR donor_profiles.area_village LIKE ?",
				searchQuery, searchQuery, searchQuery, searchQuery, searchQuery, searchQuery, searchQuery)
	}

	// Get total count for pagination
	var total int64
	query.Count(&total)

	// Fetch paginated users
	if err := query.Order("created_at desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"users": userList,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": (total + int64(limit) - 1) / int64(limit),
	})
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
	IsAdminVerified  bool       `json:"is_admin_verified"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Validate Role
	role := models.RoleDonor
	if input.Role == "admin" {
		role = models.RoleAdmin
	}

	// Generate UniqueID: BD-XXXXXX
	userID := utils.GenerateUniqueID()

	user := models.User{
		ID:           userID,
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
		UserID:           userID,
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
		IsAdminVerified:  input.IsAdminVerified,
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
	IsAvailable      *bool      `json:"is_available"`      // Pointer for false
	IsAdminVerified  *bool      `json:"is_admin_verified"` // Pointer for false
	Password         string     `json:"password"`          // Optional
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
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
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
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
		if input.IsAdminVerified != nil {
			profile.IsAdminVerified = *input.IsAdminVerified
		}
		config.DB.Save(&profile)

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
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func AdminGetDonations(c *gin.Context) {
	userID := c.Param("id")
	var donations []models.Donation
	if err := config.DB.Where("user_id = ?", userID).Order("date desc").Find(&donations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch donations"})
		return
	}
	c.JSON(http.StatusOK, donations)
}

func AdminAddDonation(c *gin.Context) {
	userID := c.Param("id")

	dateStr := c.PostForm("date")
	donationType := c.PostForm("type")
	location := c.PostForm("location")
	amountMLStr := c.PostForm("amount_ml")
	notes := c.PostForm("notes")

	if dateStr == "" || donationType == "" || location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date, Type, and Location are required"})
		return
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
	}

	var amountML int
	fmt.Sscanf(amountMLStr, "%d", &amountML)

	donation := models.Donation{
		UserID:   userID,
		Date:     date,
		Type:     donationType,
		Location: location,
		AmountML: amountML,
		Notes:    notes,
		Verified: true, // Admin added donations are auto-verified
	}

	// Handle Image Upload if present
	file, _, err := c.Request.FormFile("image")
	if err == nil {
		defer file.Close()
		ext := ".jpg"
		filename := fmt.Sprintf("donation_%d_%s%s", time.Now().Unix(), userID, ext)
		savePath, err := utils.SaveImage(file, filename, "donations")
		if err == nil {
			donation.Image = savePath
		}
	}

	if err := config.DB.Create(&donation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create donation history"})
		return
	}

	// Update last donation date in profile automatically
	var profile models.DonorProfile
	if err := config.DB.First(&profile, "user_id = ?", userID).Error; err == nil {
		// Only update if this new donation is more recent than what's stored
		if profile.LastDonationDate == nil || date.After(*profile.LastDonationDate) {
			profile.LastDonationDate = &date
			profile.IsAvailable = false
			config.DB.Save(&profile)
		}
	}

	// Send notification email
	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err == nil {
		utils.SendDonationNotification(user.Email, profile.Name, donation.Date.Format("2006-01-02"), donation.Location)
	}

	c.JSON(http.StatusCreated, donation)
}

func AdminDeleteDonation(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Donation{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete donation history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Donation record deleted"})
}

func AdminGetViewLogs(c *gin.Context) {
	userID := c.Param("id")
	var logs []models.PhoneGroupViewLog
	// Preload viewer info if needed, but for now just raw logs
	if err := config.DB.Where("target_donor_id = ?", userID).Order("created_at desc").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch view logs"})
		return
	}

	// Enrich with viewer names
	var enrichedLogs []gin.H
	for _, log := range logs {
		var viewerProfile models.DonorProfile
		config.DB.Where("user_id = ?", log.ViewerID).First(&viewerProfile)

		var viewerUser models.User
		config.DB.Where("id = ?", log.ViewerID).First(&viewerUser)

		name := viewerProfile.Name
		if name == "" {
			name = "System User"
		}

		enrichedLogs = append(enrichedLogs, gin.H{
			"id":          log.ID,
			"viewer_id":   log.ViewerID,
			"viewer_name": name,
			"unique_id":   viewerUser.ID,
			"blood_group": viewerProfile.BloodGroup,
			"phone":       viewerProfile.Phone,
			"district":    viewerProfile.District,
			"created_at":  log.CreatedAt,
			"ip_address":  log.IPAddress,
		})
	}

	c.JSON(http.StatusOK, enrichedLogs)
}
func AdminGetUserProfile(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.Where("id = ?", userID).Preload("DonorProfile").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch donations
	var donations []models.Donation
	config.DB.Where("user_id = ?", userID).Order("date desc").Find(&donations)

	// Calculate stats
	totalDonations := len(donations)
	livesSaved := totalDonations * 3
	var lastDonation *time.Time
	if totalDonations > 0 {
		lastDonation = &donations[0].Date
	}

	// Logic: Use database value strictly.
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"stats": models.DonorStats{
			TotalDonations: totalDonations,
			LivesSaved:     livesSaved,
			LastDonation:   lastDonation,
		},
		"history": donations,
	})
}

func AdminGetAllLogs(c *gin.Context) {
	// Pagination parameters
	page := 1
	limit := 20
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}
	offset := (page - 1) * limit

	// Get total count
	var total int64
	config.DB.Model(&models.PhoneGroupViewLog{}).Count(&total)

	// Fetch logs with pagination
	var logs []models.PhoneGroupViewLog
	if err := config.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}

	// Enrich with viewer and target names
	var enrichedLogs []gin.H
	for _, log := range logs {
		var viewerProfile models.DonorProfile
		config.DB.Where("user_id = ?", log.ViewerID).First(&viewerProfile)

		var viewerUser models.User
		config.DB.Where("id = ?", log.ViewerID).First(&viewerUser)

		var targetProfile models.DonorProfile
		config.DB.Where("user_id = ?", log.TargetDonorID).First(&targetProfile)

		var targetUser models.User
		config.DB.Where("id = ?", log.TargetDonorID).First(&targetUser)

		viewerName := viewerProfile.Name
		if viewerName == "" {
			viewerName = "System User"
		}

		targetName := targetProfile.Name
		if targetName == "" {
			targetName = "Unknown User"
		}

		enrichedLogs = append(enrichedLogs, gin.H{
			"id":               log.ID,
			"viewer_id":        log.ViewerID,
			"viewer_name":      viewerName,
			"viewer_unique_id": viewerUser.ID,
			"viewer_blood":     viewerProfile.BloodGroup,
			"viewer_phone":     viewerProfile.Phone,
			"viewer_email":     viewerUser.Email,
			"viewer_district":  viewerProfile.District,
			"viewer_city":      viewerProfile.City,
			"target_id":        log.TargetDonorID,
			"target_name":      targetName,
			"target_unique_id": targetUser.ID,
			"target_blood":     targetProfile.BloodGroup,
			"target_phone":     targetProfile.Phone,
			"target_email":     targetUser.Email,
			"target_district":  targetProfile.District,
			"target_city":      targetProfile.City,
			"created_at":       log.CreatedAt,
			"ip_address":       log.IPAddress,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"logs":  enrichedLogs,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func AdminGetRecentLogs(c *gin.Context) {
	var logs []models.PhoneGroupViewLog
	if err := config.DB.Order("created_at desc").Limit(10).Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recent logs"})
		return
	}

	// Enrich with viewer and target names
	var enrichedLogs []gin.H
	for _, log := range logs {
		var viewerProfile models.DonorProfile
		config.DB.Where("user_id = ?", log.ViewerID).First(&viewerProfile)

		var viewerUser models.User
		config.DB.Where("id = ?", log.ViewerID).First(&viewerUser)

		var targetProfile models.DonorProfile
		config.DB.Where("user_id = ?", log.TargetDonorID).First(&targetProfile)

		var targetUser models.User
		config.DB.Where("id = ?", log.TargetDonorID).First(&targetUser)

		viewerName := viewerProfile.Name
		if viewerName == "" {
			viewerName = "System User"
		}

		targetName := targetProfile.Name
		if targetName == "" {
			targetName = "Unknown User"
		}

		enrichedLogs = append(enrichedLogs, gin.H{
			"id":               log.ID,
			"viewer_id":        log.ViewerID,
			"viewer_name":      viewerName,
			"viewer_unique_id": viewerUser.ID,
			"viewer_blood":     viewerProfile.BloodGroup,
			"viewer_phone":     viewerProfile.Phone,
			"viewer_email":     viewerUser.Email,
			"viewer_district":  viewerProfile.District,
			"viewer_city":      viewerProfile.City,
			"target_id":        log.TargetDonorID,
			"target_name":      targetName,
			"target_unique_id": targetUser.ID,
			"target_blood":     targetProfile.BloodGroup,
			"target_phone":     targetProfile.Phone,
			"target_email":     targetUser.Email,
			"target_district":  targetProfile.District,
			"target_city":      targetProfile.City,
			"created_at":       log.CreatedAt,
			"ip_address":       log.IPAddress,
		})
	}

	c.JSON(http.StatusOK, enrichedLogs)
}

func AdminUploadProfilePicture(c *gin.Context) {
	targetUserID := c.Param("id")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}
	defer file.Close()

	if header.Size > utils.MaxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image too large (Max 5MB)"})
		return
	}

	_, err = utils.ValidateImage(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if seeker, ok := file.(io.Seeker); ok {
		seeker.Seek(0, io.SeekStart)
	}

	ext := ".jpg"
	filename := fmt.Sprintf("%s_%d_admin%s", targetUserID, time.Now().Unix(), ext)

	savePath, err := utils.SaveImage(file, filename, "profile_pictures")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	var profile models.DonorProfile
	if err := config.DB.Where("user_id = ?", targetUserID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if profile.ProfilePicture != "" {
		os.Remove(profile.ProfilePicture)
	}

	profile.ProfilePicture = savePath
	config.DB.Save(&profile)

	c.JSON(http.StatusOK, gin.H{
		"message":         "Profile picture uploaded successfully by admin",
		"profile_picture": savePath,
	})
}
