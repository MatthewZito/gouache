package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/exbotanical/gouache/models"
)

func TestHealth(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rec, req)

	var expectedBody string
	var expectedStatus int

	if name, err := os.Hostname(); err != nil {
		expectedBody = string(models.DefaultError(err.Error(), "Health check failed", 0))
		expectedStatus = http.StatusBadRequest
	} else {

		expectedBody = string(models.DefaultOk(map[string]string{"server": name}))
		expectedStatus = http.StatusOK
	}

	actualStatus := rec.Code
	actualBody := rec.Body.String()

	if actualStatus != expectedStatus {
		t.Errorf("Health returned unexpected status code: got %v want %v",
			actualStatus, http.StatusOK)
	}

	if actualBody != expectedBody {
		t.Errorf("Health returned unexpected body: got %v want %v",
			rec.Body.String(), expectedBody)
	}
}
