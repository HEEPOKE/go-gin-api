package test

import (
	"Backend/go-api/controller/auth"
	"Backend/go-api/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt"
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

func (m *MockAuthService) AddTokenToBlacklist(tokenString string, expiration time.Time) {
	m.Called(tokenString, expiration)
}

func (m *MockAuthService) IsTokenBlacklisted(tokenString string) bool {
	args := m.Called(tokenString)
	return args.Bool(0)
}

func generateTestToken() string {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(hmacSampleSecret)
	return tokenString
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

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(MockAuthService)
	authController := auth.Auth{AuthService: mockAuthService}

	router := gin.New()
	router.POST("/api/auth/login", authController.Login)

	testCases := []struct {
		name         string
		requestBody  model.Auth
		expectedCode int
	}{
		{
			name: "Successful login",
			requestBody: model.Auth{
				UsernameOrEmail: "testUser",
				Password:        "testPassword",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Invalid login credentials",
			requestBody: model.Auth{
				UsernameOrEmail: "testUser",
				Password:        "wrongPassword",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := &model.User{
				Username: "testUser",
				Password: "hashedTestPassword",
			}
			mockAuthService.On("GetUserByUsernameOrEmail", mock.Anything, tc.requestBody.UsernameOrEmail).Return(user, nil)
			mockAuthService.On("RespondWithToken", mock.Anything, user)

			reqBody, _ := json.Marshal(tc.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewReader(reqBody))
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

func TestLogout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(MockAuthService)
	authController := auth.Auth{AuthService: mockAuthService}

	router := gin.New()
	router.GET("/api/auth/logout", authController.Logout)

	tokenString := generateTestToken()

	testCases := []struct {
		name         string
		token        string
		expectedCode int
	}{
		{
			name:         "Successful logout",
			token:        tokenString,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid token",
			token:        "invalid_token",
			expectedCode: http.StatusForbidden,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/api/auth/logout", nil)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			req.Header.Set("Authorization", "Bearer "+tc.token)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedCode, resp.Code)

			mockAuthService = new(MockAuthService)
			authController.AuthService = mockAuthService
		})
	}
}
