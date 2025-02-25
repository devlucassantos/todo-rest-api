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

type MockCollectionService struct {
	mock.Mock
}

func (m *MockCollectionService) Create(collection domain.Collection, userId int) (int, error) {
	args := m.Called(collection, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockCollectionService) Update(collection domain.Collection, userId int) error {
	args := m.Called(collection, userId)
	return args.Error(0)
}

func (m *MockCollectionService) Delete(collectionId, userId int) error {
	args := m.Called(collectionId, userId)
	return args.Error(0)
}

func (m *MockCollectionService) FindAll(userId int) ([]domain.Collection, error) {
	args := m.Called(userId)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Collection), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCollection_Create(t *testing.T) {
	t.Run("should return 201 when the request is successful", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/collection", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		collectionId := 1
		mockService.On("Create", mock.Anything, mock.Anything).Return(collectionId, nil)

		_ = collectionHandler.Create(context)

		expectedBody := "{\"id\":1}\n"

		assert.Equal(t, http.StatusCreated, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user//collection", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Create(context)

		expectedBody := "{\"message\":\"Parameter not provided: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/collection",
			bytes.NewReader([]byte(invalidRequestBody)))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Create(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when name field is invalid", func(t *testing.T) {
		input := request.Collection{Name: ""}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/collection", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Create(context)

		expectedBody := "{\"message\":\"Invalid collection details.\",\"invalid_fields\":[{\"name\":\"Collection Name\"," +
			"\"description\":\"The name provided is invalid.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/collection", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		collectionId := 0
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Create", mock.Anything, mock.Anything).Return(collectionId, serviceErr)

		_ = collectionHandler.Create(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestCollection_Update(t *testing.T) {
	t.Run("should return 204 when the request is successful", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/collection/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		mockService.On("Update", mock.Anything, mock.Anything).Return(nil)

		_ = collectionHandler.Update(context)

		assert.Equal(t, http.StatusNoContent, responseData.Code)
		assert.Empty(t, responseData.Body)
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut,
			"/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/collection/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Update(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when collection ID is not a positive integer", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/collection/-10",
			bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "-10")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Update(context)

		expectedBody := "{\"message\":\"Invalid parameter: Collection ID\",\"invalid_fields\":[{" +
			"\"name\":\"Collection ID\",\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/collection/2",
			bytes.NewBuffer([]byte(invalidRequestBody)))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Update(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when name field is invalid", func(t *testing.T) {
		input := request.Collection{Name: ""}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/collection/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Update(context)

		expectedBody := "{\"message\":\"Invalid collection details.\",\"invalid_fields\":[{\"name\":\"Collection Name\"," +
			"\"description\":\"The name provided is invalid.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Collection{Name: "Test collection"}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/collection/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Update", mock.Anything, mock.Anything).Return(serviceErr)

		_ = collectionHandler.Update(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestCollection_Delete(t *testing.T) {
	t.Run("should return 204 when the request is successful", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/collection/2", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		mockService.On("Delete", mock.Anything, mock.Anything).Return(nil)

		_ = collectionHandler.Delete(context)

		assert.Equal(t, http.StatusNoContent, responseData.Code)
		assert.Empty(t, responseData.Body)
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete,
			"/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/collection/2", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Delete(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when collection ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/collection/-10",
			nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "-10")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.Delete(context)

		expectedBody := "{\"message\":\"Invalid parameter: Collection ID\",\"invalid_fields\":[{" +
			"\"name\":\"Collection ID\",\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/collection/2", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Delete", mock.Anything, mock.Anything).Return(serviceErr)

		_ = collectionHandler.Delete(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestCollection_FindAll(t *testing.T) {
	t.Run("should return 200 when the request is successful", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/collection", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		collections := []domain.Collection{
			*domain.NewCollection(1, "Test Collection 1"),
			*domain.NewCollection(2, "Test Collection 2"),
		}
		mockService.On("FindAll", mock.Anything).Return(collections, nil)

		_ = collectionHandler.FindAll(context)

		expectedBody := "[{\"id\":1,\"name\":\"Test Collection 1\"},{\"id\":2,\"name\":\"Test Collection 2\"}]\n"

		assert.Equal(t, http.StatusOK, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/collection",
			nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}

		_ = collectionHandler.FindAll(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/collection", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockCollectionService)
		collectionHandler := Collection{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("FindAll", mock.Anything).Return(nil, serviceErr)

		_ = collectionHandler.FindAll(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}
