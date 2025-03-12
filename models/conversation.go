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

type ConversationUser struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ConversationID uuid.UUID  `gorm:"type:uuid;not null;index"`
	MemberID       uuid.UUID  `gorm:"type:uuid;not null;index"` // User in this conversation
	CreatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	Member       User         `gorm:"foreignKey:MemberID"`
}
