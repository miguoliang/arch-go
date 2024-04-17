package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type ListResponse[T any] struct {
	Items []T `json:"items,omitempty"`
}
