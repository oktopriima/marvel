package dto

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	Duration     int64  `json:"duration"`
	RefreshToken string `json:"refresh_token"`
}
