package main

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load .env file from the root of backend (assuming running from backend dir)
	if err := godotenv.Load(); err != nil {
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
			},
			Profile: models.DonorProfile{
				Name:        "Rahim Uddin",
				BloodGroup:  "A+",
				Phone:       "01711111111",
				District:    "Dhaka",
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
				Phone:       "01822222222",
				District:    "Chittagong",
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
				Phone:       "01933333333",
				District:    "Dhaka",
				City:        "Dhaka",
				AreaVillage: "Dhanmondi",
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
				Phone:       "01644444444",
				District:    "Sylhet",
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
				Phone:       "01555555555",
				District:    "Khulna",
				City:        "Khulna",
				AreaVillage: "Sonadanga",
				IsAvailable: true,
			},
		},
	}

	fmt.Println("Seeding database...")

	for _, u := range users {
		var existingUser models.User
		if err := db.Where("email = ?", u.User.Email).First(&existingUser).Error; err == nil {
			fmt.Printf("User %s already exists, skipping.\n", u.User.Email)
			// Assuming if user exists, we might want to ensure profile/donations exist or just skip.
			// For simplicity, skipping entire block if user exists.
			// In a real seed, `FirstOrCreate` is better.
			continue
		}

		tx := db.Begin()
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
	}

	fmt.Println("Seeding complete.")
}

func parseTime(value string) time.Time {
	t, _ := time.Parse("2006-01-02", value)
	return t
}
