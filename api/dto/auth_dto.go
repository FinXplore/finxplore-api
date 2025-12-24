package dto

// Request
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	DialingCode string `json:"dialing_code" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
	Currency string `json:"currency"`
}

// Response
type UserResponse struct {
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
}

type LoginRequest struct {
	DialingCode string `json:"dialing_code" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
	UserCode     string `json:"user_code"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	DialingCode string `json:"dialing_code" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"` // Changed from user_code
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}