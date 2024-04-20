package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreatedResponse struct {
	Id string `json:"id"`
}
