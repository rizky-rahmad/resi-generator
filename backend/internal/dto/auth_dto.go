package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	ExpiresIn int    `json:"expires_in"`
}
