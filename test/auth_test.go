package test

import (
	AuthController "Backend/go-api/controller/auth"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/api/auth/register", AuthController.Register)

	testCases := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Successful registration",
			requestBody:  `{"username":"testUser","password":"testPassword","email":"test@email.com","tel":"123456789"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "User already exists",
			requestBody:  `{"username":"existingUser","password":"testPassword","email":"test@email.com","tel":"123456789"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid JSON",
			requestBody:  `{"username":"invalidJson","password":"testPassword",}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(tc.requestBody))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedCode, resp.Code)
		})
	}
}
