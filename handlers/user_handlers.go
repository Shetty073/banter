package handlers

import (
	"banter/models"
	"banter/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserDetailsHandler fetches user details by ID
// @Summary Get User Details
// @Description Fetches user details by user ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 404 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /user/{id} [get]
// @Security AuthorizationToken
func GetUserDetailsHandler(c *gin.Context) {
	// Get user ID from URL parameter
	idParam := c.Param("id")

	// Convert ID to UUID
	userID, err := uuid.Parse(idParam)
	if err != nil {
		responses.BadRequest(c, "Invalid User ID", "User ID must be a valid UUID")
		return
	}

	// Fetch user by ID
	user, err := models.GetUserByID(userID)
	if err != nil {
		responses.NotFound(c, "User Not Found", "No user found with the given ID")
		return
	}

	// Respond with user details
	responses.Ok(c, gin.H{
		"id":            user.ID,
		"username":      user.Username,
		"email":         user.Email,
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"date_of_birth": user.DateOfBirth,
		"gender":        user.Gender,
		"mobile_number": user.MobileNumber,
		"profile_photo": user.ProfilePhoto,
		"is_staff":      user.IsStaff,
		"is_owner":      user.IsOwner,
		"last_seen":     user.LastSeen,
		"status":        user.Status,
		"created_at":    user.CreatedAt,
		"updated_at":    user.UpdatedAt,
	})
}
