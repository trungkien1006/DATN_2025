package authdto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"superadmin@example.com"`
	Password string `json:"password" validate:"required,password_strong" example:"Admin1234"`
	// Password string `json:"password" validate:"required" example:"Abc12345"`

	Device string `json:"device" validate:"required,oneof=mobile web" example:"web"`
}

type GetAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
