package models

import (
	"banter/constants/enums"
	"banter/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username        string           `gorm:"unique;not null;index"`
	Email           string           `gorm:"unique;not null;index"`
	Password        string           `gorm:"not null"`
	FirstName       string           `gorm:"type:varchar(50)"`
	LastName        string           `gorm:"type:varchar(50)"`
	DateOfBirth     time.Time        `gorm:"type:date"`
	Gender          string           `gorm:"type:varchar(12)"`
	MobileNumber    string           `gorm:"type:varchar(15);index"`
	ProfilePhoto    string           `gorm:"type:varchar(1024);"`
	ProfilePhotoUrl string           `gorm:"type:varchar(1024);"`
	IsStaff         bool             `gorm:"default:false;index"`
	IsOwner         bool             `gorm:"default:false;index"`
	LastSeen        *time.Time       `gorm:"type:timestamp;index"`
	Status          enums.UserStatus `gorm:"type:varchar(15);index"`
	CreatedAt       *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt       *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt       gorm.DeletedAt
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	if err := stores.GetDb().First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (u *User) CreateUser() error {
	u.ID = uuid.New()
	return stores.GetDb().Create(u).Error
}

// UpdateUser updates an existing user
func (u *User) UpdateUser() error {
	return stores.GetDb().Save(u).Error
}

// DeleteUser marks the user as inactive
func (u *User) DeleteUser() error {
	return stores.GetDb().Delete(u).Error
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]User, error) {
	var users []User
	if err := stores.GetDb().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmailOrUsername retrieves a user by email or username
func GetUserByEmailOrUsername(email, username string, user *User) error {
	db := stores.GetDb()

	if email != "" {
		err := db.Where("email = ?", email).First(user).Error
		if err == nil {
			return nil
		}
	}

	if username != "" {
		err := db.Where("username = ?", username).First(user).Error
		if err == nil {
			return nil
		}
	}

	return gorm.ErrRecordNotFound
}
