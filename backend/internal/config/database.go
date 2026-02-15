package config

import (
	"fmt"
	"log"
	"os"

	"blood-donor-system/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully")

	// Auto Migrate
	err = DB.AutoMigrate(
		&models.User{},
		&models.DonorProfile{},
		&models.SearchLog{},
		&models.PhoneGroupViewLog{},
		&models.Donation{},
		&models.LocationRegistry{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	SeedLocationRegistry()
}

func SeedLocationRegistry() {
	var profiles []models.DonorProfile
	DB.Find(&profiles)

	for _, p := range profiles {
		if p.AreaVillage != "" {
			var existing models.LocationRegistry
			if err := DB.Where("name = ? AND type = ? AND district = ?", p.AreaVillage, "village", p.District).First(&existing).Error; err != nil {
				DB.Create(&models.LocationRegistry{
					Name:     p.AreaVillage,
					Type:     "village",
					District: p.District,
					Upazila:  p.Upazila,
				})
			}
		}
		if p.City != "" {
			var existing models.LocationRegistry
			if err := DB.Where("name = ? AND type = ? AND district = ?", p.City, "city", p.District).First(&existing).Error; err != nil {
				DB.Create(&models.LocationRegistry{
					Name:     p.City,
					Type:     "city",
					District: p.District,
					Upazila:  p.Upazila,
				})
			}
		}
	}
}
