package models

import (
	"time"

	"gorm.io/gorm"
)

type Attachment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	MessageID uint           `gorm:"not null" json:"message_id"`
	FilePath  string         `gorm:"type:varchar(2048)"`
	FileType  string         `gorm:"type:varchar(10)"` // image, video or document
	FileUrl   string         `gorm:"type:varchar(1024)"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
