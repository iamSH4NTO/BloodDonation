package handlers

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"blood-donor-system/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateProfileInput struct {
	Name             string       `json:"name" binding:"required"`
	Gender           string       `json:"gender"`
	Birthday         *time.Time   `json:"birthday"` // Added Birthday
	BloodGroup       string       `json:"blood_group" binding:"required"`
	District         string       `json:"district" binding:"required"`
	City             string       `json:"city" binding:"required"`
	AreaVillage      string       `json:"area_village"`
	PostalCode       string       `json:"postal_code"`
	Latitude         float64      `json:"latitude"`
	Longitude        float64      `json:"longitude"`
	GoogleMapLink    string       `json:"google_map_link"`
	LastDonationDate *time.Time   `json:"last_donation_date"`
	IsAvailable      bool         `json:"is_available"`
	PrivacySettings  models.JSONB `json:"privacy_settings"`
}

func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profile models.DonorProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Update fields
	profile.Name = input.Name
	profile.Gender = input.Gender
	profile.Birthday = input.Birthday // Save Birthday
	profile.BloodGroup = input.BloodGroup
	profile.District = input.District
	profile.City = input.City
	profile.AreaVillage = input.AreaVillage
	profile.PostalCode = input.PostalCode
	profile.Latitude = input.Latitude
	profile.Longitude = input.Longitude
	profile.GoogleMapLink = input.GoogleMapLink
	profile.LastDonationDate = input.LastDonationDate
	profile.IsAvailable = input.IsAvailable

	if input.PrivacySettings != nil {
		profile.PrivacySettings = input.PrivacySettings
	}

	// Transaction to update profile and locations
	tx := config.DB.Begin()
	if err := tx.Save(&profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	// Register locations in Registry
	if profile.AreaVillage != "" {
		var existing models.LocationRegistry
		if err := tx.Where("name = ? AND type = ? AND district = ?", profile.AreaVillage, "village", profile.District).First(&existing).Error; err != nil {
			tx.Create(&models.LocationRegistry{
				Name:     profile.AreaVillage,
				Type:     "village",
				District: profile.District,
				Upazila:  profile.Upazila,
			})
		}
	}
	if profile.City != "" {
		var existing models.LocationRegistry
		if err := tx.Where("name = ? AND type = ? AND district = ?", profile.City, "city", profile.District).First(&existing).Error; err != nil {
			tx.Create(&models.LocationRegistry{
				Name:     profile.City,
				Type:     "city",
				District: profile.District,
				Upazila:  profile.Upazila,
			})
		}
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated", "profile": profile})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var profile models.DonorProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Fetch donations
	var donations []models.Donation
	config.DB.Where("user_id = ?", userID).Order("date desc").Find(&donations)

	// Calculate stats
	totalDonations := len(donations)
	livesSaved := totalDonations * 3 // Rough estimate
	var lastDonation *time.Time
	if totalDonations > 0 {
		lastDonation = &donations[0].Date
	}

	// Logic: Use database value strictly. No more dynamic auto-available logic.
	response := models.ProfileResponse{
		Profile: profile,
		Stats: models.DonorStats{
			TotalDonations: totalDonations,
			LivesSaved:     livesSaved,
			LastDonation:   lastDonation,
		},
		History: donations,
	}

	c.JSON(http.StatusOK, response)
}

func GetDonors(c *gin.Context) {
	group := c.Query("group")
	district := c.Query("district")
	upazila := c.Query("upazila")
	gender := c.Query("gender")
	q := c.Query("q")
	// Pagination
	// lat/long for distance sort (Phase 3)

	// Validate required params
	if group == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blood group is required"})
		return
	}

	var donors []models.DonorProfile
	query := config.DB.Model(&models.DonorProfile{})

	// Check for availability filter (default to true if not specified "false")
	availableOnly := c.Query("available_only")

	if availableOnly == "true" {
		// Strict Logic: Only show those marked available in DB
		query = query.Where("is_available = ?", true)
	}

	query = query.Where("blood_group = ?", group)

	if district != "" {
		query = query.Where("district = ?", district)
	}
	if upazila != "" {
		query = query.Where("upazila = ?", upazila)
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if q != "" {
		query = query.Where("area_village LIKE ? OR city LIKE ? OR district LIKE ? OR name LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%", "%"+q+"%")
	}

	// Pagination
	var total int64
	query.Count(&total)

	page := 1
	limit := 12 // Default match frontend itemsPerPage
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 12
	}

	offset := (page - 1) * limit
	query.Offset(offset).Limit(limit).Find(&donors)

	// Minimize response to public fields only (Privacy)
	type DonorSearchDTO struct {
		UserID           string       `json:"user_id"`
		Name             string       `json:"name"`
		Gender           string       `json:"gender"`
		BloodGroup       string       `json:"blood_group"`
		Division         string       `json:"division"`
		District         string       `json:"district"`
		Upazila          string       `json:"upazila"`
		City             string       `json:"city"`
		AreaVillage      string       `json:"area_village"`
		LastDonationDate *time.Time   `json:"last_donation_date"`
		IsAvailable      bool         `json:"is_available"`
		PrivacySettings  models.JSONB `json:"privacy_settings"`
	}

	searchShowResults := make([]DonorSearchDTO, len(donors))

	// Map results to DTO strictly from DB status
	for i := range donors {
		searchShowResults[i] = DonorSearchDTO{
			UserID:           donors[i].UserID,
			Name:             donors[i].Name,
			Gender:           donors[i].Gender,
			BloodGroup:       donors[i].BloodGroup,
			Division:         donors[i].Division,
			District:         donors[i].District,
			Upazila:          donors[i].Upazila,
			City:             donors[i].City,
			AreaVillage:      donors[i].AreaVillage,
			LastDonationDate: donors[i].LastDonationDate,
			IsAvailable:      donors[i].IsAvailable, // Strict DB value
			PrivacySettings:  donors[i].PrivacySettings,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"donors": searchShowResults,
		"total":  total,
		"page":   page,
		"limit":  limit,
	})
}

func SearchLocations(c *gin.Context) {
	q := c.Query("q")
	if len(q) < 2 {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	searchQuery := "%" + q + "%"
	var locations []models.LocationRegistry
	config.DB.Where("name LIKE ? OR district LIKE ? OR upazila LIKE ?", searchQuery, searchQuery, searchQuery).
		Limit(10).
		Find(&locations)

	var results []gin.H
	for _, loc := range locations {
		results = append(results, gin.H{
			"name":     loc.Name,
			"type":     loc.Type,
			"subtitle": loc.Upazila + ", " + loc.District,
			"district": loc.District,
			"upazila":  loc.Upazila,
		})
	}

	c.JSON(http.StatusOK, results)
}

func GetDonor(c *gin.Context) {
	id := c.Param("id")
	var donor models.DonorProfile
	if err := config.DB.First(&donor, "user_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	// Fetch donations for public profile
	var donations []models.Donation
	config.DB.Where("user_id = ?", donor.UserID).Order("date desc").Find(&donations)

	// Calculate stats
	totalDonations := len(donations)
	livesSaved := totalDonations * 3
	var lastDonation *time.Time
	if totalDonations > 0 {
		lastDonation = &donations[0].Date
	}

	// Logic: Use database value strictly.
	response := models.ProfileResponse{
		Profile: donor,
		Stats: models.DonorStats{
			TotalDonations: totalDonations,
			LivesSaved:     livesSaved,
			LastDonation:   lastDonation,
		},
		History: donations,
	}

	c.JSON(http.StatusOK, response)
}

func GetDonorContact(c *gin.Context) {
	donorID := c.Param("id")
	viewerID := c.MustGet("userID").(string)

	var donor models.DonorProfile
	if err := config.DB.Where("user_id = ?", donorID).First(&donor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	// Log the view
	viewLog := models.PhoneGroupViewLog{
		ViewerID:      viewerID,
		TargetDonorID: donor.UserID,
		IPAddress:     c.ClientIP(),
	}

	config.DB.Create(&viewLog)

	c.JSON(http.StatusOK, gin.H{"phone": donor.Phone})
}

type AddDonationInput struct {
	Date     time.Time `json:"date" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	Location string    `json:"location" binding:"required"`
	AmountML int       `json:"amount_ml"`
	Notes    string    `json:"notes"`
}

func AddDonation(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var input AddDonationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donation := models.Donation{
		UserID:   userID,
		Date:     input.Date,
		Type:     input.Type,
		Location: input.Location,
		AmountML: input.AmountML,
		Notes:    input.Notes,
		Verified: false, // Default to unverified for self-reported
	}

	if err := config.DB.Create(&donation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save donation"})
		return
	}

	// Update last donation date in profile automatically
	var profile models.DonorProfile
	if err := config.DB.First(&profile, "user_id = ?", userID).Error; err == nil {
		// Only update if this new donation is more recent than what's stored
		if profile.LastDonationDate == nil || input.Date.After(*profile.LastDonationDate) {
			profile.LastDonationDate = &input.Date
			profile.IsAvailable = false
			config.DB.Save(&profile)
		}
	}

	// Send notification email
	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err == nil {
		utils.SendDonationNotification(user.Email, profile.Name, donation.Date.Format("2006-01-02"), donation.Location)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Donation added successfully", "donation": donation})
}
