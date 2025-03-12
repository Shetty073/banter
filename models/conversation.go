package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(50)"` // Group name (if applicable)
	IsGroup   bool       `gorm:"default:false;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
}
