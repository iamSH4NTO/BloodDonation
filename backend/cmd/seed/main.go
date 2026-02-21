package main

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"blood-donor-system/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load .env file (try current dir and parent dirs)
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../../.env")
	}
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Connect to Database
	config.ConnectDB()
	db := config.DB

	// Default password for all users
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	users := []struct {
		User      models.User
		Profile   models.DonorProfile
		Donations []models.Donation
	}{
		{
			User: models.User{
				Email:        "admin@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleAdmin,
				IsActive:     true,
				IsVerified:   true,
			},
			Profile: models.DonorProfile{
				Name:  "System Admin",
				Phone: "01700000000",
			},
		},
		{
			User: models.User{
				Email:        "donor1@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
				IsVerified:   true,
			},
			Profile: models.DonorProfile{
				Name:        "Rahim Uddin",
				BloodGroup:  "A+",
				Gender:      "Male",
				Birthday:    getPtr(parseTime("1995-05-15")), // ~29 years old
				Phone:       "01711111111",
				District:    "Dhaka",
				Upazila:     "Mirpur",
				City:        "Dhaka",
				AreaVillage: "Mirpur 10",
				IsAvailable: true,
			},
			Donations: []models.Donation{
				{
					Type:     "Whole Blood",
					Location: "Dhaka Medical College",
					AmountML: 450,
					Verified: true,
					Date:     parseTime("2023-01-15"),
				},
				{
					Type:     "Platelets",
					Location: "Square Hospital",
					AmountML: 250,
					Verified: true,
					Date:     parseTime("2023-05-20"),
				},
			},
		},
		{
			User: models.User{
				Email:        "donor2@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
			},
			Profile: models.DonorProfile{
				Name:        "Karim Ahmed",
				BloodGroup:  "B-",
				Gender:      "Male",
				Birthday:    getPtr(parseTime("1990-10-20")), // ~34 years old
				Phone:       "01822222222",
				District:    "Chittagong",
				Upazila:     "Patiya",
				City:        "Chittagong",
				AreaVillage: "Halishahar",
				IsAvailable: true,
			},
			Donations: []models.Donation{
				{
					Type:     "Whole Blood",
					Location: "Chittagong Medical College",
					AmountML: 450,
					Verified: true,
					Date:     parseTime("2023-08-10"),
				},
			},
		},
		{
			User: models.User{
				Email:        "donor3@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
			},
			Profile: models.DonorProfile{
				Name:        "Fatima Begum",
				BloodGroup:  "O+",
				Gender:      "Female",
				Birthday:    getPtr(parseTime("1998-03-12")), // ~26 years old
				Phone:       "01933333333",
				District:    "Dhaka",
				Upazila:     "Dhanmondi",
				City:        "Dhaka",
				AreaVillage: "Road 10",
				IsAvailable: true,
			},
			Donations: []models.Donation{
				{
					Type:     "Whole Blood",
					Location: "Red Crescent Society",
					AmountML: 450,
					Verified: true,
					Date:     parseTime("2022-11-05"),
				},
				{
					Type:     "Whole Blood",
					Location: "City Hospital",
					AmountML: 450,
					Verified: true,
					Date:     parseTime("2023-03-12"),
				},
				{
					Type:     "Whole Blood",
					Location: "Dhaka Medical College",
					AmountML: 450,
					Verified: true,
					Date:     parseTime("2023-07-25"),
				},
			},
		},
		{
			User: models.User{
				Email:        "donor4@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
			},
			Profile: models.DonorProfile{
				Name:        "Sujon Miah",
				BloodGroup:  "AB+",
				Gender:      "Male",
				Birthday:    getPtr(parseTime("2000-01-01")), // ~24 years old
				Phone:       "01644444444",
				District:    "Sylhet",
				Upazila:     "Sylhet Sadar", // Assuming Sylhet Sadar exists or generic
				City:        "Sylhet",
				AreaVillage: "Zindabazar",
				IsAvailable: true,
			},
		},
		{
			User: models.User{
				Email:        "donor5@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
			},
			Profile: models.DonorProfile{
				Name:        "Nusrat Jahan",
				BloodGroup:  "A-",
				Gender:      "Female",
				Birthday:    getPtr(parseTime("1992-07-07")), // ~32 years old
				Phone:       "01555555555",
				District:    "Gazipur",
				Upazila:     "Gazipur Sadar",
				City:        "Gazipur",
				AreaVillage: "Joydebpur",
				IsAvailable: true,
			},
		},
		{
			User: models.User{
				Email:        "donor6@example.com",
				PasswordHash: string(hashedPassword),
				Role:         models.RoleDonor,
				IsActive:     true,
			},
			Profile: models.DonorProfile{
				Name:        "Mehedi Hasan",
				BloodGroup:  "O-",
				Gender:      "Male",
				Birthday:    getPtr(parseTime("1996-12-12")),
				Phone:       "01712345678",
				District:    "Dhaka",
				Upazila:     "Savar",
				City:        "Dhaka",
				AreaVillage: "Hemayetpur",
				IsAvailable: true,
			},
		},
	}

	fmt.Println("Seeding database...")

	for _, u := range users {
		var existingUser models.User
		if err := db.Where("email = ?", u.User.Email).First(&existingUser).Error; err == nil {
			fmt.Printf("User %s exists, updating profile...\n", u.User.Email)

			// Update User verification status
			db.Model(&models.User{}).Where("id = ?", existingUser.ID).Updates(models.User{IsVerified: u.User.IsVerified})

			var profile models.DonorProfile
			if err := db.Where("user_id = ?", existingUser.ID).First(&profile).Error; err == nil {
				// Update fields
				updateData := map[string]interface{}{
					"gender":   u.Profile.Gender,
					"birthday": u.Profile.Birthday,
					"upazila":  u.Profile.Upazila,
				}
				if err := db.Model(&models.DonorProfile{}).Where("user_id = ?", profile.UserID).Updates(updateData).Error; err != nil {
					log.Printf("Failed to update profile for %s: %v\n", u.User.Email, err)
				}
			}
			continue
		}

		tx := db.Begin()
		u.User.ID = utils.GenerateUniqueID()
		if err := tx.Create(&u.User).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create user %s: %v\n", u.User.Email, err)
			continue
		}

		u.Profile.UserID = u.User.ID
		if err := tx.Create(&u.Profile).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create profile for %s: %v\n", u.User.Email, err)
			continue
		}

		for _, d := range u.Donations {
			d.UserID = u.User.ID
			if err := tx.Create(&d).Error; err != nil {
				tx.Rollback()
				log.Printf("Failed to create donation for %s: %v\n", u.User.Email, err)
				continue
			}
		}

		tx.Commit()
		fmt.Printf("Created user: %s with %d donations\n", u.User.Email, len(u.Donations))

		// Small sleep to ensure uniqueness of BD-XXXXXX IDs (based on UnixNano)
		time.Sleep(2 * time.Millisecond)
	}

	fmt.Println("Seeding complete.")
}

func parseTime(value string) time.Time {
	t, _ := time.Parse("2006-01-02", value)
	return t
}

func getPtr(t time.Time) *time.Time {
	return &t
}
