package types

type RegisterUserRequest struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"email"`
}

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"email"`
}

type TokenResponse struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
