package models

import (
	"banter/stores"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Conversation model represents a chat group or direct message.
type Conversation struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string         `gorm:"type:varchar(50)"`
	IsGroup        bool           `gorm:"default:false;index"`
	GroupPhotoPath string         `gorm:"type:varchar(1024);"`
	GroupPhotoUrl  string         `gorm:"type:varchar(1024);"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// ConversationMember represents the members in a conversation.
type ConversationMember struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ConversationID uuid.UUID      `gorm:"type:uuid;not null;index"`
	MemberID       uuid.UUID      `gorm:"type:uuid;not null;index"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	Member       User         `gorm:"foreignKey:MemberID"`
}

type ConversationWithMembers struct {
	Conversation *Conversation `json:"conversation"`
	Members      []*User       `json:"members"`
}

type PaginatedConversations struct {
	Conversations []ConversationWithMembers `json:"conversations"`
	CurrentPage   int                       `json:"current_page"`
	NextPage      int                       `json:"next_page,omitempty"`
	HasNextPage   bool                      `json:"has_next_page"`
}

// CreateConversation inserts a new conversation into the database.
func (c *Conversation) CreateConversation() error {
	return stores.GetDb().Create(c).Error
}

// GetConversationByID fetches a conversation by its ID along with its members.
func GetConversationByID(id uuid.UUID) (*ConversationWithMembers, error) {
	var conversation Conversation
	var result ConversationWithMembers
	var members []*User

	result = ConversationWithMembers{
		Conversation: nil,
		Members:      nil,
	}

	// Fetch conversation data
	err := stores.GetDb().
		Joins("JOIN conversation_members ON conversations.id = conversation_members.conversation_id").
		Where("conversations.deleted_at is null AND conversations.id = ?", id).
		Find(&conversation).Error
	if err != nil {
		return &ConversationWithMembers{}, err
	}

	if conversation.ID.String() != "00000000-0000-0000-0000-000000000000" {
		err = stores.GetDb().
			Joins("JOIN conversation_members ON users.id = conversation_members.member_id").
			Where("conversation_members.conversation_id = ?", id).
			Find(&members).Error
		if err != nil {
			return &ConversationWithMembers{}, err
		}

		result = ConversationWithMembers{
			Conversation: &conversation,
			Members:      members,
		}
	}

	return &result, nil
}

// UpdateConversation updates the details of an existing conversation.
func (c *Conversation) UpdateConversation() error {
	return stores.GetDb().Save(c).Error
}

// DeleteConversation deletes a conversation (soft delete by default).
func DeleteConversation(id uuid.UUID) (error, error) {
	err1 := stores.GetDb().Where("id = ?", id).Delete(&Conversation{}).Error
	err2 := stores.GetDb().Where("conversation_id = ?", id).Delete(&ConversationMember{}).Error

	return err1, err2
}

// AddMembers adds multiple users to a conversation.
func AddMembers(conversationID uuid.UUID, memberIDs []uuid.UUID) error {
	var members []ConversationMember
	for _, memberID := range memberIDs {
		members = append(members, ConversationMember{
			ID:             uuid.New(),
			ConversationID: conversationID,
			MemberID:       memberID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}
	return stores.GetDb().Create(&members).Error
}

// GetMembers fetches all members of a conversation.
func GetMembers(conversationID uuid.UUID) ([]ConversationMember, error) {
	var members []ConversationMember
	if err := stores.GetDb().Where("conversation_id = ?", conversationID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// RemoveMember removes a user from a conversation.
// RemoveMember removes a user from a conversation only if more than 2 members exist.
func RemoveMember(conversationID, memberID uuid.UUID) error {
	db := stores.GetDb()
	var memberCount int64

	// Count the number of members in the conversation
	err := db.Model(&ConversationMember{}).
		Where("conversation_id = ?", conversationID).
		Count(&memberCount).Error

	if err != nil {
		return err
	}

	// Prevent removal if only 2 members are left
	if memberCount <= 2 {
		return errors.New("cannot remove member, only 2 members left")
	}

	// Remove the member
	return db.Where("conversation_id = ? AND member_id = ?", conversationID, memberID).
		Delete(&ConversationMember{}).Error
}

// GetUserConversations fetches paginated conversations for a given user along with member details.
func GetUserConversations(userID uuid.UUID, page, limit int) (PaginatedConversations, error) {
	var conversations []Conversation
	var results []ConversationWithMembers
	var totalConversations int64
	offset := (page - 1) * limit

	// Count total conversations the user is part of
	err := stores.GetDb().
		Table("conversations").
		Joins("JOIN conversation_members ON conversations.id = conversation_members.conversation_id").
		Where("conversations.deleted_at is null AND conversation_members.member_id = ?", userID).
		Count(&totalConversations).Error
	if err != nil {
		return PaginatedConversations{}, err
	}

	// Get paginated conversations
	err = stores.GetDb().
		Joins("JOIN conversation_members ON conversations.id = conversation_members.conversation_id").
		Where("conversation_members.member_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&conversations).Error
	if err != nil {
		return PaginatedConversations{}, err
	}

	// Fetch members for each conversation
	for _, conversation := range conversations {
		var members []*User
		err := stores.GetDb().
			Joins("JOIN conversation_members ON users.id = conversation_members.member_id").
			Where("conversation_members.conversation_id = ?", conversation.ID).
			Find(&members).Error
		if err != nil {
			return PaginatedConversations{}, err
		}

		results = append(results, ConversationWithMembers{
			Conversation: &conversation,
			Members:      members,
		})
	}

	// Determine next page
	totalPages := int((totalConversations + int64(limit) - 1) / int64(limit)) // Ceiling division
	hasNextPage := page < totalPages
	nextPage := page + 1
	if !hasNextPage {
		nextPage = 0 // No next page
	}

	return PaginatedConversations{
		Conversations: results,
		CurrentPage:   page,
		NextPage:      nextPage,
		HasNextPage:   hasNextPage,
	}, nil
}

// GetAllConversations fetches all conversations.
func GetAllConversations() ([]Conversation, error) {
	var conversations []Conversation
	if err := stores.GetDb().Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

// RestoreConversation restores a soft-deleted conversation.
func RestoreConversation(id uuid.UUID) error {
	return stores.GetDb().Model(&Conversation{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}
