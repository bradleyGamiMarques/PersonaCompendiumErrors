package personacompendiumerrors

import (
	"time"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
}

func CreateErrorResponse(err string, message string, path string, statusCode int) ErrorResponse {
	return ErrorResponse{
		Error:      err,
		Message:    message,
		Path:       path,
		StatusCode: statusCode,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}
