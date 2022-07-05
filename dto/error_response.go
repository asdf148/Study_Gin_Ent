package dto

type ErrorResponse interface {}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(error_message string) ErrorResponse {
	return &errorResponse{
		ErrorMessage: error_message,
	}
}