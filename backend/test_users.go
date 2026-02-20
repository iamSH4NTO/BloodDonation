package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	IsVerified   bool
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/blood_donor?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Default to root:root@tcp(127.0.0.1:3306) if env vars are empty
	if os.Getenv("DB_USER") == "" {
		dsn = "root:root@tcp(127.0.0.1:3306)/blood_donor?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	db.Find(&users)
	for _, u := range users {
		fmt.Printf("Email: %s, Verified: %v, ID: %s\n", u.Email, u.IsVerified, u.ID)

		// Auto-verify if not verified
		if !u.IsVerified {
			db.Model(&u).Update("is_verified", true)
			fmt.Println("Auto-verified", u.Email)
		}
	}
}
