package personacompendiumerrors_test

import (
	"testing"
	"time"

	personacompendiumerrors "github.com/bradleyGamiMarques/PersonaCompendiumErrors"
)

func TestCreateErrorResponse(t *testing.T) {
	errResp := personacompendiumerrors.CreateErrorResponse("Test Error", "This is a test error message", "/test/path", 400)

	if errResp.Error != "Test Error" {
		t.Errorf("expected Error to be 'Test Error', got '%s'", errResp.Error)
	}
	if errResp.Message != "This is a test error message" {
		t.Errorf("expected Message to be 'This is a test error message', got '%s'", errResp.Message)
	}
	if errResp.Path != "/test/path" {
		t.Errorf("expected Path to be '/test/path', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 400 {
		t.Errorf("expected StatusCode to be 400, got %d", errResp.StatusCode)
	}
	if _, err := time.Parse(time.RFC3339, errResp.Timestamp); err != nil {
		t.Errorf("expected Timestamp to be in RFC3339 format, got '%s'", errResp.Timestamp)
	}
}

func TestBadRequestError(t *testing.T) {
	errResp := personacompendiumerrors.BadRequestError("This is a bad request", "/bad/request")

	if errResp.Error != "Bad Request" {
		t.Errorf("expected Error to be 'Bad Request', got '%s'", errResp.Error)
	}
	if errResp.Message != "This is a bad request" {
		t.Errorf("expected Message to be 'This is a bad request', got '%s'", errResp.Message)
	}
	if errResp.Path != "/bad/request" {
		t.Errorf("expected Path to be '/bad/request', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 400 {
		t.Errorf("expected StatusCode to be 400, got %d", errResp.StatusCode)
	}
}

func TestUnauthorizedError(t *testing.T) {
	errResp := personacompendiumerrors.UnauthorizedError("This is an unauthorized request", "/unauthorized")

	if errResp.Error != "Unauthorized" {
		t.Errorf("expected Error to be 'Unauthorized', got '%s'", errResp.Error)
	}
	if errResp.Message != "This is an unauthorized request" {
		t.Errorf("expected Message to be 'This is an unauthorized request', got '%s'", errResp.Message)
	}
	if errResp.Path != "/unauthorized" {
		t.Errorf("expected Path to be '/unauthorized', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 401 {
		t.Errorf("expected StatusCode to be 401, got %d", errResp.StatusCode)
	}
}

func TestNotForbiddenError(t *testing.T) {
	errResp := personacompendiumerrors.ForbiddenError("This is a forbidden request", "/forbidden")

	if errResp.Error != "Forbidden" {
		t.Errorf("expected Error to be 'Forbidden', got '%s'", errResp.Error)
	}
	if errResp.Message != "This is a forbidden request" {
		t.Errorf("expected Message to be 'This is a forbidden request', got '%s'", errResp.Message)
	}
	if errResp.Path != "/forbidden" {
		t.Errorf("expected Path to be '/forbidden', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 403 {
		t.Errorf("expected StatusCode to be 403, got %d", errResp.StatusCode)
	}
}

func TestNotFoundError(t *testing.T) {
	errResp := personacompendiumerrors.NotFoundError("This resource was not found", "/not/found")

	if errResp.Error != "Not Found" {
		t.Errorf("expected Error to be 'Not Found', got '%s'", errResp.Error)
	}
	if errResp.Message != "This resource was not found" {
		t.Errorf("expected Message to be 'This resource was not found', got '%s'", errResp.Message)
	}
	if errResp.Path != "/not/found" {
		t.Errorf("expected Path to be '/not/found', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 404 {
		t.Errorf("expected StatusCode to be 404, got %d", errResp.StatusCode)
	}
}

func TestInternalServerError(t *testing.T) {
	errResp := personacompendiumerrors.InternalServerError("Internal Server Error", "/error")

	if errResp.Error != "Internal Server Error" {
		t.Errorf("expected Error to be 'Internal Server Error', got '%s'", errResp.Error)
	}
	if errResp.Message != "Internal Server Error" {
		t.Errorf("expected Message to be 'Internal Server Error', got '%s'", errResp.Message)
	}
	if errResp.Path != "/error" {
		t.Errorf("expected Path to be '/error', got '%s'", errResp.Path)
	}
	if errResp.StatusCode != 500 {
		t.Errorf("expected StatusCode to be 500, got %d", errResp.StatusCode)
	}
}

func TestJSONResponse(t *testing.T) {
	errResp := personacompendiumerrors.InternalServerError("This is an internal error", "/internal/error")
	response, err := personacompendiumerrors.JSONResponse(errResp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.StatusCode != 500 {
		t.Errorf("expected StatusCode to be 500, got %d", response.StatusCode)
	}
	if response.Headers["Content-Type"] != "application/json" {
		t.Errorf("expected Content-Type to be 'application/json', got '%s'", response.Headers["Content-Type"])
	}
	if response.Headers["Access-Control-Allow-Origin"] != "*" {
		t.Errorf("expected Access-Control-Allow-Origin to be '*', got '%s'", response.Headers["Access-Control-Allow-Origin"])
	}
	if response.Body == "" {
		t.Errorf("expected Body to be non-empty")
	}
}
