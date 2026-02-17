package main

import (
	"blood-donor-system/internal/config"
	"blood-donor-system/internal/handlers"
	"blood-donor-system/internal/middleware"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load() // Ignore error if .env not found (e.g. prod env vars)

	config.ConnectDB()
	config.LoadSMTPConfig()

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware())

	// Auth Routes
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/verify-email", handlers.VerifyEmail)
		auth.POST("/forgot-password", handlers.ForgotPassword)
		auth.POST("/reset-password", handlers.ResetPassword)
	}

	// Protected Route Example
	r.GET("/api/v1/ping", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"user_id": c.MustGet("userID"),
		})
	})

	donors := r.Group("/api/v1/donors")
	{
		donors.GET("", handlers.GetDonors)
		donors.GET("/locations/search", handlers.SearchLocations)
		donors.GET("/:id", handlers.GetDonor)
		donors.GET("/:id/contact", middleware.AuthMiddleware(), handlers.GetDonorContact)
	}

	// Profile routes
	r.GET("/api/v1/profile", middleware.AuthMiddleware(), handlers.GetProfile)
	r.PUT("/api/v1/profile", middleware.AuthMiddleware(), handlers.UpdateProfile)
	r.POST("/api/v1/donations", middleware.AuthMiddleware(), handlers.AddDonation)

	// Admin Routes
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/stats", handlers.GetDashboardStats)
		admin.GET("/users", handlers.GetAllUsers)
		admin.POST("/users", handlers.CreateUser)
		admin.GET("/users/:id", handlers.AdminGetUserProfile)
		admin.PUT("/users/:id", handlers.UpdateUser)
		admin.DELETE("/users/:id", handlers.DeleteUser)
		admin.GET("/users/:id/donations", handlers.AdminGetDonations)
		admin.POST("/users/:id/donations", handlers.AdminAddDonation)
		admin.DELETE("/donations/:id", handlers.AdminDeleteDonation)
		admin.GET("/users/:id/view-logs", handlers.AdminGetViewLogs)
		admin.GET("/logs", handlers.AdminGetAllLogs)
		admin.GET("/logs/recent", handlers.AdminGetRecentLogs)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
