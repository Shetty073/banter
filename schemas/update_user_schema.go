package schemas

type UpdateUserSchema struct {
	Username     *string `json:"username" binding:"omitempty,alphanum"`
	Email        *string `json:"email" binding:"omitempty,email"`
	Password     *string `json:"password" binding:"omitempty,min=8,max=32"`
	FirstName    *string `json:"first_name" binding:"omitempty,min=2,max=50,alpha"`
	LastName     *string `json:"last_name" binding:"omitempty,min=2,max=50,alpha"`
	DateOfBirth  *string `json:"date_of_birth" binding:"omitempty"` // Keep as string
	Gender       *string `json:"gender" binding:"omitempty,oneof=male female other"`
	MobileNumber *string `json:"mobile_number" binding:"omitempty,len=10,numeric"`
}
