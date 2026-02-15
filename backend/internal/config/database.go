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
		&models.PhoneGroupViewLog{},
		&models.Donation{},
		&models.LocationRegistry{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Seed locations
	SeedLocationRegistry()

	// Ensure all existing numeric IDs are migrated to BD-XXXXXX format
	MigrateUniqueIDs()
}

func MigrateUniqueIDs() {
	var users []models.User
	// Fetch all users whose IDs don't start with "BD-"
	DB.Where("id NOT LIKE ?", "BD-%").Find(&users)

	if len(users) == 0 {
		return
	}

	log.Printf("Found %d users to migrate to Unique ID format", len(users))

	for _, user := range users {
		var oldIDInt int
		fmt.Sscanf(user.ID, "%d", &oldIDInt)
		if oldIDInt == 0 {
			continue
		}

		newID := fmt.Sprintf("BD-%06d", (oldIDInt*1000+123456)%1000000)
		log.Printf("Migrating user %s -> %s", user.ID, newID)

		// Update User ID and all foreign keys manually in a transaction
		tx := DB.Begin()
		tx.Exec("SET FOREIGN_KEY_CHECKS = 0")

		tx.Model(&models.DonorProfile{}).Where("user_id = ?", user.ID).Update("user_id", newID)
		tx.Model(&models.Donation{}).Where("user_id = ?", user.ID).Update("user_id", newID)
		tx.Model(&models.PhoneGroupViewLog{}).Where("viewer_id = ?", user.ID).Update("viewer_id", newID)
		tx.Model(&models.PhoneGroupViewLog{}).Where("target_donor_id = ?", user.ID).Update("target_donor_id", newID)

		if err := tx.Exec("UPDATE users SET id = ? WHERE id = ?", newID, user.ID).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to migrate user %s: %v", user.ID, err)
			continue
		}

		tx.Exec("SET FOREIGN_KEY_CHECKS = 1")
		tx.Commit()
	}
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
