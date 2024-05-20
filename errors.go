package personacompendiumerrors

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
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

func BadRequestError(message string, path string) ErrorResponse {
	return CreateErrorResponse("Bad Request", message, path, 400)
}

func UnauthorizedError(message string, path string) ErrorResponse {
	return CreateErrorResponse("Unauthorized", message, path, 401)
}

func ForbiddenError(message string, path string) ErrorResponse {
	return CreateErrorResponse("Forbidden", message, path, 403)
}

func NotFoundError(message string, path string) ErrorResponse {
	return CreateErrorResponse("Not Found", message, path, 404)
}

func InternalServerError(message string, path string) ErrorResponse {
	return CreateErrorResponse("Internal Server Error", message, path, 500)
}

func (errResp ErrorResponse) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(errResp)
	if err != nil {
		return "", nil
	}
	return string(jsonBytes), nil
}

func JSONResponse(errResp ErrorResponse) (events.APIGatewayProxyResponse, error) {
	errorResponseJSON, err := errResp.ToJSON()
	if err != nil {
		// Return a generic error response if marshaling fails
		genericError := ErrorResponse{
			Error:      "Internal Server Error",
			Message:    "An error occurred while processing the error response",
			Path:       errResp.Path,
			StatusCode: 500,
			Timestamp:  time.Now().Format((time.RFC3339)),
		}
		genericErrorJSON, _ := json.Marshal(genericError)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			MultiValueHeaders: nil,
			Body:              string(genericErrorJSON),
			IsBase64Encoded:   false,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode:        errResp.StatusCode,
		Headers:           map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		MultiValueHeaders: nil,
		Body:              errorResponseJSON,
		IsBase64Encoded:   false,
	}, nil
}
