package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thefonzie-codes/goLift/internal/testutils"
)

func TestHealthCheck(t *testing.T) {
	stats := testutils.NewTestStats("handlers")
	stats.LogInfo(t, "Starting health check test...")

	// Create a new handler
	stats.LogInfo(t, "Creating new handler")
	h := NewHandler(nil)

	// Create a test request
	stats.LogInfo(t, "Creating test request")
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		stats.LogError(t, "Failed to create request")
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	stats.LogInfo(t, "Creating response recorder")
	rr := httptest.NewRecorder()

	// Call the handler
	stats.LogInfo(t, "Calling health check handler")
	h.HealthCheck(rr, req)

	// Check the status code
	if status := rr.Code; status == http.StatusOK {
		stats.LogSuccess(t, "Status code is 200 OK")
	} else {
		stats.LogError(t, "Wrong status code")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	stats.LogInfo(t, "Checking response body")
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		stats.LogError(t, "Failed to decode response body")
		t.Fatal(err)
	}

	expected := "ok"
	if response["status"] == expected {
		stats.LogSuccess(t, "Response body contains correct status")
	} else {
		stats.LogError(t, "Response body contains incorrect status")
		t.Errorf("handler returned unexpected body: got %v want %v",
			response["status"], expected)
	}

	stats.PrintSummary(t)
}
