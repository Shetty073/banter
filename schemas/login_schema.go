package schemas

// LoginSchema holds the structure for user login data
type LoginSchema struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}
