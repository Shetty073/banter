package handlers

import (
	"banter/models"
	"banter/responses"
	"banter/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// StartConversationHandler starts a new chat or group conversation
// @Summary Start Conversation
// @Description Creates a new direct or group chat
// @Tags Conversation
// @Accept json
// @Produce json
// @Param conversation body schemas.StartConversationSchema true "Conversation Data"
// @Success 201 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversation [post]
// @Security AuthorizationToken
func StartConversationHandler(c *gin.Context) {
	var input schemas.StartConversationSchema

	// Parse request body
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.BadRequest(c, "Invalid Input", err.Error())
		return
	}

	// Ensure there are at least 3 members in a group chat
	if input.IsGroup && len(input.Members) < 3 {
		responses.BadRequest(c, "Invalid Members", "Group chats must have at least 3 members")
		return
	}

	// Generate conversation ID
	conversationID := uuid.New()

	// Create new conversation
	conversation := models.Conversation{
		ID:      conversationID,
		Name:    input.Name,
		IsGroup: input.IsGroup,
	}

	// Save conversation to DB
	if err := conversation.CreateConversation(); err != nil {
		responses.InternalServerError(c, "Failed to create conversation", err.Error())
		return
	}

	// Add members to conversation
	if err := models.AddMembers(conversationID, input.Members); err != nil {
		responses.InternalServerError(c, "Failed to add members", err.Error())
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message":       "Conversation created successfully",
		"conversation":  conversation,
		"members_count": len(input.Members),
	})
}

// GetConversationsHandler fetches all conversations for a user with pagination
// @Summary Get User Conversations
// @Tags Conversation
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Success 200 {array} models.ConversationWithMembers
// @Failure 400 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversations/member/{user_id} [get]
// @Security AuthorizationToken
func GetConversationsHandler(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		responses.BadRequest(c, "Invalid User ID", "Must be a valid UUID")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	conversations, err := models.GetUserConversations(userID, page, limit)
	if err != nil {
		responses.InternalServerError(c, "Failed to fetch conversations", err.Error())
		return
	}

	c.JSON(http.StatusOK, conversations)
}

// GetConversationHandler fetches conversation details
// @Summary Get Conversation Details
// @Tags Conversation
// @Accept json
// @Produce json
// @Param id path string true "Conversation ID"
// @Success 200 {object} models.Conversation
// @Failure 400 {object} responses.FailureBody
// @Failure 404 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversation/{id} [get]
// @Security AuthorizationToken
func GetConversationHandler(c *gin.Context) {
	conversationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		responses.BadRequest(c, "Invalid Conversation ID", "Must be a valid UUID")
		return
	}

	conversation, err := models.GetConversationByID(conversationID)
	if err != nil {
		responses.NotFound(c, "Conversation Not Found", "No conversation found with the given ID")
		return
	}

	c.JSON(http.StatusOK, conversation)
}

// AddMemberHandler adds a user to a conversation
// @Summary Add Member
// @Tags Conversation
// @Accept json
// @Produce json
// @Param id path string true "Conversation ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversation/{id}/member/{user_id} [post]
// @Security AuthorizationToken
func AddMemberHandler(c *gin.Context) {
	conversationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		responses.BadRequest(c, "Invalid Conversation ID", "Must be a valid UUID")
		return
	}

	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		responses.BadRequest(c, "Invalid User ID", "Must be a valid UUID")
		return
	}

	if err := models.AddMembers(conversationID, []uuid.UUID{userID}); err != nil {
		responses.InternalServerError(c, "Failed to add member", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}

// RemoveMemberHandler removes a user from a conversation
// @Summary Remove Member
// @Tags Conversation
// @Accept json
// @Produce json
// @Param id path string true "Conversation ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversation/{id}/member/{user_id} [delete]
// @Security AuthorizationToken
func RemoveMemberHandler(c *gin.Context) {
	conversationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		responses.BadRequest(c, "Invalid Conversation ID", "Must be a valid UUID")
		return
	}
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		responses.BadRequest(c, "Invalid User ID", "Must be a valid UUID")
		return
	}

	if err := models.RemoveMember(conversationID, userID); err != nil {
		responses.InternalServerError(c, "Failed to remove member", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}

// DeleteConversationHandler deletes a conversation
// @Summary Delete Conversation
// @Tags Conversation
// @Accept json
// @Produce json
// @Param id path string true "Conversation ID"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /conversation/{id} [delete]
// @Security AuthorizationToken
func DeleteConversationHandler(c *gin.Context) {
	conversationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		responses.BadRequest(c, "Invalid Conversation ID", "Must be a valid UUID")
		return
	}

	if err1, err2 := models.DeleteConversation(conversationID); err1 != nil && err2 != nil {
		responses.InternalServerError(c, "Failed to delete conversation", err1.Error()+err2.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation deleted successfully"})
}
