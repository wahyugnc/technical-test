package response

type UserDataResponse struct {
	Data  []User `json:"data"`
	Total int    `json:"total"`
}
