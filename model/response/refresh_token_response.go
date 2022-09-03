package response

type RefreshTokenResponse struct {
	StatusCode  int    `json:"statusCode"`
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}
