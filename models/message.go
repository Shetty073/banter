package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ConversationID uuid.UUID  `gorm:"type:uuid;not null;index"`
	SenderID       uuid.UUID  `gorm:"type:uuid;not null;index"`
	Content        string     `gorm:"type:varchar(1024)"`
	AttachmentURL  string     `gorm:"type:varchar(50)"`
	AttachmentType string     `gorm:"type:varchar(10)"` // image, video or document
	CreatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	Sender       User         `gorm:"foreignKey:SenderID"`
}
