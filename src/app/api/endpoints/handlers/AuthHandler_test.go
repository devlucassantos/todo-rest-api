package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/src/app/api/endpoints/dto/request"
	"todo/src/core/domain"
	"todo/src/core/projecterrors/todoerrors"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) SignUp(account domain.Account) (*domain.Account, error) {
	args := m.Called(account)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Account), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthService) SignIn(account domain.Account) (*domain.Account, error) {
	args := m.Called(account)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Account), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAuth_SignUp(t *testing.T) {
	t.Run("should return 201 when the request is successful", func(t *testing.T) {
		input := request.Account{Name: "John Doe", Email: "johndoe@example.com", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}
		accountDomain := domain.NewAccount(1, input.Name, input.Email, input.Password, "")
		mockService.On("SignUp", mock.Anything).Return(accountDomain, nil)

		_ = authHandler.SignUp(context)

		expectedBody := "{\"id\":1,\"name\":\"John Doe\",\"email\":\"johndoe@example.com\",\"access_token\":\"\"}\n"

		assert.Equal(t, http.StatusCreated, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signup",
			bytes.NewReader([]byte(invalidRequestBody)))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignUp(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when email field is invalid", func(t *testing.T) {
		input := request.Account{Name: "John Doe", Email: "invalid-email", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewReader(requestBody))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignUp(context)

		expectedBody := "{\"message\":\"Invalid account details.\",\"invalid_fields\":[{\"name\":\"Account Email\"," +
			"\"description\":\"The email provided is invalid.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when password field is invalid", func(t *testing.T) {
		input := request.Account{Name: "John Doe", Email: "johndoe@example.com", Password: "pass"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewReader(requestBody))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignUp(context)

		expectedBody := "{\"message\":\"Invalid account details.\",\"invalid_fields\":[{\"name\":\"Account Password\"," +
			"\"description\":\"The password provided is invalid. The password must be between 8 and 50 characters.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Account{Name: "John Doe", Email: "johndoe@example.com", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("SignUp", mock.Anything).Return(nil, serviceErr)

		_ = authHandler.SignUp(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestAuth_SignIn(t *testing.T) {
	t.Run("should return 200 when the request is successful", func(t *testing.T) {
		input := request.Account{Email: "johndoe@example.com", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}
		accountDomain := domain.NewAccount(1, "", input.Email, input.Password, "")
		mockService.On("SignIn", mock.Anything).Return(accountDomain, nil)

		_ = authHandler.SignIn(context)

		expectedBody := "{\"id\":1,\"name\":\"\",\"email\":\"johndoe@example.com\",\"access_token\":\"\"}\n"

		assert.Equal(t, http.StatusOK, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signin",
			bytes.NewReader([]byte(invalidRequestBody)))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignIn(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when email field is invalid", func(t *testing.T) {
		input := request.Account{Email: "invalid-email", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signin", bytes.NewReader(requestBody))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignIn(context)

		expectedBody := "{\"message\":\"Invalid account details.\",\"invalid_fields\":[{\"name\":\"Account Email\"," +
			"\"description\":\"The email provided is invalid.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when password field is invalid", func(t *testing.T) {
		input := request.Account{Email: "johndoe@example.com", Password: "pass"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signin", bytes.NewReader(requestBody))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}

		_ = authHandler.SignIn(context)

		expectedBody := "{\"message\":\"Invalid account details.\",\"invalid_fields\":[{\"name\":\"Account Password\"," +
			"\"description\":\"The password provided is invalid. The password must be between 8 and 50 characters.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Account{Email: "johndoe@example.com", Password: "SecurePass123!"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/auth/signin", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockAuthService)
		authHandler := Auth{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("SignIn", mock.Anything).Return(nil, serviceErr)

		_ = authHandler.SignIn(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}
