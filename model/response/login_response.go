package response

type LoginResponse struct {
	StatusCode   int    `json:"statusCode"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
