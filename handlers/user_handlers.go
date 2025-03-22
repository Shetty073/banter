package handlers

import (
	"banter/models"
	"banter/responses"
	"banter/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
		"profile_photo": user.ProfilePhotoUrl,
		"is_staff":      user.IsStaff,
		"is_owner":      user.IsOwner,
		"last_seen":     user.LastSeen,
		"status":        user.Status,
		"created_at":    user.CreatedAt,
		"updated_at":    user.UpdatedAt,
	})
}

// UpdateUserDetailsHandler updates user details by ID
// @Summary Update User Details
// @Description Updates user details by user ID. Only the provided fields will be updated.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body schemas.UpdateUserSchema true "User update data"
// @Success 200 {object} responses.SuccessBody "User details updated successfully"
// @Failure 400 {object} responses.FailureBody "Invalid request data"
// @Failure 404 {object} responses.FailureBody "User not found"
// @Failure 500 {object} responses.FailureBody "Internal server error"
// @Router /user/{id} [patch]
// @Security AuthorizationToken
func UpdateUserDetailsHandler(c *gin.Context) {
	// Get user ID from URL parameter
	idParam := c.Param("id")

	// Convert ID to UUID
	userID, err := uuid.Parse(idParam)
	if err != nil {
		responses.BadRequest(c, "Invalid User ID", "User ID must be a valid UUID")
		return
	}

	var updatedUserDataInput schemas.UpdateUserSchema

	// Bind JSON request body to input struct
	if err := c.ShouldBindJSON(&updatedUserDataInput); err != nil {
		// Improved error handling to show validation failures
		responses.BadRequest(c, "Invalid Input", err.Error())
		return
	}

	// Fetch user by ID
	user, err := models.GetUserByID(userID)
	if err != nil {
		responses.NotFound(c, "User Not Found", "No user found with the given ID")
		return
	}

	// Update only the provided fields
	if updatedUserDataInput.Username != nil {
		user.Username = *updatedUserDataInput.Username
	}
	if updatedUserDataInput.Email != nil {
		user.Email = *updatedUserDataInput.Email
	}
	if updatedUserDataInput.Password != nil {
		// Proceed with password hashing
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*updatedUserDataInput.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.InternalServerError(c, "Failed to hash password", err.Error())
			return
		}
		user.Password = string(hashedPassword)
	}
	if updatedUserDataInput.FirstName != nil {
		user.FirstName = *updatedUserDataInput.FirstName
	}
	if updatedUserDataInput.LastName != nil {
		user.LastName = *updatedUserDataInput.LastName
	}
	if updatedUserDataInput.DateOfBirth != nil {
		dob, err := schemas.ParseDOB(*updatedUserDataInput.DateOfBirth)
		if err != nil {
			responses.BadRequest(c, "Invalid Input", err.Error())
			return
		}
		user.DateOfBirth = dob
	}
	if updatedUserDataInput.Gender != nil {
		user.Gender = *updatedUserDataInput.Gender
	}
	if updatedUserDataInput.MobileNumber != nil {
		user.MobileNumber = *updatedUserDataInput.MobileNumber
	}

	// Save updated user data
	err = user.UpdateUser()
	if err != nil {
		responses.InternalServerError(c, "Data Updation Error", "Error updating data")
		return
	}

	err = user.UpdateUser()
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
		"profile_photo": user.ProfilePhotoUrl,
		"is_staff":      user.IsStaff,
		"is_owner":      user.IsOwner,
		"last_seen":     user.LastSeen,
		"status":        user.Status,
		"created_at":    user.CreatedAt,
		"updated_at":    user.UpdatedAt,
	})
}
