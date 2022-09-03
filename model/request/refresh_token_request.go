package request

type RefreshTokenRequest struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refreshToken"`
}
