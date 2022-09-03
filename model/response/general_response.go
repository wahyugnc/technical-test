package response

type GeneralResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
