package schemas

import (
	"github.com/google/uuid"
)

type StartConversationSchema struct {
	Name    string      `json:"name" binding:"omitempty,min=1,max=100"`
	IsGroup bool        `json:"is_group" binding:"omitempty"`
	Members []uuid.UUID `json:"members" binding:"required,min=1,max=1024,dive,required"`
}
