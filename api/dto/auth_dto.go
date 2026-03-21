package dto

// Request
type RegisterRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name"`
	DialingCode  string `json:"dialing_code"`
	MobileNumber string `json:"mobile_number"`
	Currency     string `json:"currency"`
}

// Response
type UserResponse struct {
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
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
	Email        string `json:"email" binding:"required,email"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
	Currency  string `json:"currency"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
}

type ProfileResponse struct {
	UserCode      string `json:"user_code"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AvatarURL     string `json:"avatar_url"`
	Currency      string `json:"currency"`
	Role          string `json:"role"`
	RiskTolerance string `json:"risk_tolerance"`
}