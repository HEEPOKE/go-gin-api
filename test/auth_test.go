package test

import (
	"Backend/go-api/controller/auth"
	"Backend/go-api/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) RegisterUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthService) CheckUserExistence(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *MockAuthService) GetUserByUsernameOrEmail(c *gin.Context, usernameOrEmail string) (*model.User, error) {
	args := m.Called(c, usernameOrEmail)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockAuthService) RespondWithToken(c *gin.Context, user *model.User) {
	m.Called(c, user)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(MockAuthService)
	authController := auth.Auth{AuthService: mockAuthService}

	router := gin.New()
	router.POST("/api/auth/register", authController.Register)

	testCases := []struct {
		name         string
		requestBody  model.User
		expectedCode int
	}{
		{
			name: "Successful registration",
			requestBody: model.User{
				Username: "testUser",
				Password: "testPassword",
				Email:    "test@email.com",
				Tel:      "123456789",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "User already exists",
			requestBody: model.User{
				Username: "existingUser",
				Password: "testPassword",
				Email:    "test@email.com",
				Tel:      "123456789",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuthService.On("CheckUserExistence", tc.requestBody.Username).Return(false, nil)
			mockAuthService.On("RegisterUser", mock.Anything).Return(nil)

			reqBody, _ := json.Marshal(tc.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewReader(reqBody))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedCode, resp.Code)

			mockAuthService = new(MockAuthService)
			authController.AuthService = mockAuthService
		})
	}
}
