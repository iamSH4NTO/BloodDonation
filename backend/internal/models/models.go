package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleDonor   Role = "donor"
)

type User struct {
	ID                string         `gorm:"primaryKey;size:20" json:"id"` // BD-123456
	Email             string         `gorm:"uniqueIndex;not null;size:191" json:"email"`
	PasswordHash      string         `gorm:"not null" json:"-"`
	Role              Role           `gorm:"type:enum('admin','manager','donor');default:'donor'" json:"role"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	IsVerified        bool           `gorm:"default:false" json:"is_verified"`
	VerificationToken string         `gorm:"size:100" json:"-"`
	ResetToken        string         `gorm:"size:100" json:"-"`
	ResetTokenExpires *time.Time     `json:"-"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	DonorProfile      *DonorProfile  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"donor_profile"`
}

type DonorProfile struct {
	UserID           string     `gorm:"primaryKey;size:20" json:"user_id"` // Belongs to User
	Name             string     `gorm:"not null" json:"name"`
	Gender           string     `gorm:"size:10" json:"gender"`           // Male, Female, Other
	Birthday         *time.Time `json:"birthday"`                        // Added Birthday
	BloodGroup       string     `gorm:"index;size:5" json:"blood_group"` // A+, B-, etc.
	Phone            string     `gorm:"not null" json:"phone"`
	Division         string     `gorm:"index" json:"division"` // Added Division
	District         string     `gorm:"index" json:"district"`
	Upazila          string     `gorm:"index" json:"upazila"` // Added Upazila
	City             string     `gorm:"index" json:"city"`
	AreaVillage      string     `json:"area_village"`
	PostalCode       string     `json:"postal_code"`
	Latitude         float64    `gorm:"default:0" json:"latitude"`
	Longitude        float64    `gorm:"default:0" json:"longitude"`
	GoogleMapLink    string     `json:"google_map_link"`
	LastDonationDate *time.Time `json:"last_donation_date"`
	IsAvailable      bool       `gorm:"default:true" json:"is_available"`
	IsAdminVerified  bool       `gorm:"default:false" json:"is_admin_verified"` // Admin manually verified the donor
	PrivacySettings  JSONB      `gorm:"type:json" json:"privacy_settings"`      // Custom type or use string for MySQL JSON
	ProfilePicture   string     `json:"profile_picture"`
	FacebookLink     string     `json:"facebook_link"`
	InstagramLink    string     `json:"instagram_link"`
	LinkedinLink     string     `json:"linkedin_link"`
	YoutubeLink      string     `json:"youtube_link"`
}

// JSONB is a helper for JSON fields if needed, or simple string for MySQL JSON
type JSONB []byte

type PhoneGroupViewLog struct {
	ID            uint   `gorm:"primaryKey"`
	ViewerID      string `gorm:"index;size:20"`
	TargetDonorID string `gorm:"index;size:20"`
	CreatedAt     time.Time
	IPAddress     string
}

// Donation represents a single donation event
type Donation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"index;size:20" json:"user_id"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"` // Whole Blood, Platelets, etc.
	Location  string    `json:"location"`
	AmountML  int       `json:"amount_ml"`
	Notes     string    `json:"notes"` // Optional details
	Image     string    `json:"image"` // EVIDENCE IMAGE
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}

type LocationRegistry struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;size:191" json:"name"` // Village or City name
	Type      string `gorm:"size:20" json:"type"`        // "village" or "city"
	District  string `gorm:"size:100" json:"district"`
	Upazila   string `gorm:"size:100" json:"upazila"`
	Division  string `gorm:"size:100" json:"division"`
	CreatedAt time.Time
}

// ProfileResponse bundles profile and stats
type ProfileResponse struct {
	Profile DonorProfile `json:"profile"`
	Stats   DonorStats   `json:"stats"`
	History []Donation   `json:"history"`
}

type DonorStats struct {
	TotalDonations int        `json:"total_donations"`
	LivesSaved     int        `json:"lives_saved"` // Calculated
	LastDonation   *time.Time `json:"last_donation"`
	CurrentBadge   string     `json:"current_badge"` // Added for Gamification
	NextBadgeAt    int        `json:"next_badge_at"` // Added for Gamification
}
