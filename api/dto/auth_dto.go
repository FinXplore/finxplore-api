package dto

// Request
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	DialingCode string `json:"dialing_code" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
}

// Response
type UserResponse struct {
	// ID        string `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
}