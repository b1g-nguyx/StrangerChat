package httpdelivery

// AdminLoginRequest defines the payload structure for admin login.
// Notice it might have extra security fields like OTP.
type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	OTP      string `json:"otp" validate:"required,len=6"` // Example: Admin requires 2FA
}

// AdminAuthData defines the response data for admin login.
// Notice it might return different fields than the regular Client login.
type AdminAuthData struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}
