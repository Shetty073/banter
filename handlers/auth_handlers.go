package handlers

import (
	"banter/models"
	"banter/responses"
	"banter/schemas"
	"banter/utils/config"
	"banter/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles customer registration
// @Summary Register a new customer
// @Description Creates a new user with the provided details
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schemas.RegisterSchema true "User registration data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var input schemas.RegisterSchema

	// Bind JSON request body to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		// Improved error handling to show validation failures
		responses.BadRequest(c, "Invalid Input", err.Error())
		return
	}

	// Proceed with password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.InternalServerError(c, "Hashing Error", "Failed to hash password")
		return
	}

	dob, err := input.GetParsedDOB()
	if err != nil {
		responses.BadRequest(c, "Invalid Input", err.Error())
		return
	}

	// Create user model
	user := models.User{
		ID:           uuid.New(),
		Username:     input.Username,
		Email:        input.Email,
		Password:     string(hashedPassword),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		DateOfBirth:  dob,
		Gender:       input.Gender,
		MobileNumber: input.MobileNumber,
	}

	// Save user to the database
	if err := user.CreateUser(); err != nil {
		responses.InternalServerError(c, err.Error(), "Failed to create user")
		return
	}

	// Success response
	responses.Created(c, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles user authentication
// @Summary User login
// @Description Authenticates a user with email/username and password, and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schemas.LoginSchema true "User login data"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 401 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var input schemas.LoginSchema

	// Bind JSON request body to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		// Improved error handling to show validation failures
		responses.BadRequest(c, "Invalid Input", err.Error())
		return
	}

	if input.Email == "" && input.Username == "" {
		responses.BadRequest(c, "Invalid Input", "Please provide username or email")
		return
	}

	// Find user by email or username
	var user models.User
	err := models.GetUserByEmailOrUsername(input.Email, input.Username, &user)
	if err != nil {
		responses.Unauthorized(c, "Authentication Error", err.Error())
		return
	}

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		responses.Unauthorized(c, "Authentication Error", err.Error())
		return
	}

	// Generate JWT token using your existing function
	tokenString, err := jwt.GenerateToken(user.ID.String(), config.Configs.Auth.TokenValidityInHrs)
	if err != nil {
		responses.InternalServerError(c, "Token Generation Error", "Failed to generate token")
		return
	}

	// Respond with token
	responses.Ok(c, gin.H{"token": tokenString})
}
