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

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) Create(task domain.Task, userId int) (int, error) {
	args := m.Called(task, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskService) Update(task domain.Task, userId int) error {
	args := m.Called(task, userId)
	return args.Error(0)
}

func (m *MockTaskService) Delete(taskId, userId int) error {
	args := m.Called(taskId, userId)
	return args.Error(0)
}

func (m *MockTaskService) FindAll(userId int) ([]domain.Task, error) {
	args := m.Called(userId)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskService) FindByCollectionId(collectionId, userId int) ([]domain.Task, error) {
	args := m.Called(collectionId, userId)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestTask_Create(t *testing.T) {
	t.Run("should return 201 when the request is successful", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/task", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		taskId := 1
		mockService.On("Create", mock.Anything, mock.Anything).Return(taskId, nil)

		_ = taskHandler.Create(context)

		expectedBody := "{\"id\":1}\n"

		assert.Equal(t, http.StatusCreated, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user//task", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Create(context)

		expectedBody := "{\"message\":\"Parameter not provided: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/task",
			bytes.NewReader([]byte(invalidRequestBody)))
		requestData.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Create(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPost, "/user/1/task", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		taskId := 0
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Create", mock.Anything, mock.Anything).Return(taskId, serviceErr)

		_ = taskHandler.Create(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestTask_Update(t *testing.T) {
	t.Run("should return 204 when the request is successful", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/task/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		mockService.On("Update", mock.Anything, mock.Anything).Return(nil)

		_ = taskHandler.Update(context)

		assert.Equal(t, http.StatusNoContent, responseData.Code)
		assert.Empty(t, responseData.Body)
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/task/2",
			bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Update(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when task ID is not a positive integer", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/task/-10",
			bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "-10")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Update(context)

		expectedBody := "{\"message\":\"Invalid parameter: Task ID\",\"invalid_fields\":[{" +
			"\"name\":\"Task ID\",\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 400 when request body is not a valid JSON", func(t *testing.T) {
		invalidRequestBody := "{invalid_json}"
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/task/2",
			bytes.NewBuffer([]byte(invalidRequestBody)))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Update(context)

		expectedBody := "{\"message\":\"The request format is invalid.\"}\n"

		assert.Equal(t, http.StatusBadRequest, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		input := request.Task{Description: "Task Description", Finished: false, CollectionId: 1}
		requestBody, _ := json.Marshal(input)
		requestData := httptest.NewRequest(http.MethodPut, "/user/1/task/2", bytes.NewBuffer(requestBody))
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Update", mock.Anything, mock.Anything).Return(serviceErr)

		_ = taskHandler.Update(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestTask_Delete(t *testing.T) {
	t.Run("should return 204 when the request is successful", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/task/2", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		mockService.On("Delete", mock.Anything, mock.Anything).Return(nil)

		_ = taskHandler.Delete(context)

		assert.Equal(t, http.StatusNoContent, responseData.Code)
		assert.Empty(t, responseData.Body)
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/task/2",
			nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Delete(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when task ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/task/-10",
			nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "-10")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.Delete(context)

		expectedBody := "{\"message\":\"Invalid parameter: Task ID\",\"invalid_fields\":[{" +
			"\"name\":\"Task ID\",\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodDelete, "/user/1/task/2", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "taskId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("Delete", mock.Anything, mock.Anything).Return(serviceErr)

		_ = taskHandler.Delete(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestTask_FindAll(t *testing.T) {
	t.Run("should return 200 when the request is successful", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		tasks := []domain.Task{
			*domain.NewTask(1, "Test Task 1", true,
				domain.NewCollection(2, "Test Collection 2")),
			*domain.NewTask(2, "Test Task 2", false,
				domain.NewCollection(1, "Test Collection 1")),
		}
		mockService.On("FindAll", mock.Anything).Return(tasks, nil)

		_ = taskHandler.FindAll(context)

		expectedBody := "[{\"id\":1,\"description\":\"Test Task 1\",\"finished\":true,\"collection\":{\"id\":2," +
			"\"name\":\"Test Collection 2\"}},{\"id\":2,\"description\":\"Test Task 2\",\"finished\":false," +
			"\"collection\":{\"id\":1,\"name\":\"Test Collection 1\"}}]\n"

		assert.Equal(t, http.StatusOK, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/task",
			nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.FindAll(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId")
		context.SetParamValues("1")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("FindAll", mock.Anything).Return(nil, serviceErr)

		_ = taskHandler.FindAll(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}

func TestTask_FindByCollectionId(t *testing.T) {
	t.Run("should return 200 when the request is successful", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/collection/2/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		tasks := []domain.Task{
			*domain.NewTask(1, "Test Task 1", true,
				domain.NewCollection(2, "Test Collection 2")),
			*domain.NewTask(2, "Test Task 2", false,
				domain.NewCollection(2, "Test Collection 2")),
		}
		mockService.On("FindByCollectionId", mock.Anything, mock.Anything).Return(tasks, nil)

		_ = taskHandler.FindByCollectionId(context)

		expectedBody := "[{\"id\":1,\"description\":\"Test Task 1\",\"finished\":true,\"collection\":{\"id\":2," +
			"\"name\":\"Test Collection 2\"}},{\"id\":2,\"description\":\"Test Task 2\",\"finished\":false," +
			"\"collection\":{\"id\":2,\"name\":\"Test Collection 2\"}}]\n"

		assert.Equal(t, http.StatusOK, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when user ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet,
			"/user/4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa/collection/2/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("4f626b9f-ee9a-4e41-a6a6-44d4833dfdfa", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.FindByCollectionId(context)

		expectedBody := "{\"message\":\"Invalid parameter: User ID\",\"invalid_fields\":[{\"name\":\"User ID\"," +
			"\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 422 when collection ID is not a positive integer", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/collection/-10/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "-10")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}

		_ = taskHandler.FindByCollectionId(context)

		expectedBody := "{\"message\":\"Invalid parameter: Collection ID\",\"invalid_fields\":[{" +
			"\"name\":\"Collection ID\",\"description\":\"Conversion error.\"}]}\n"

		assert.Equal(t, http.StatusUnprocessableEntity, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})

	t.Run("should return 500 when a service layer error is returned", func(t *testing.T) {
		requestData := httptest.NewRequest(http.MethodGet, "/user/1/collection/2/task", nil)
		requestData.Header.Set("Content-Type", "application/json")
		responseData := httptest.NewRecorder()
		context := echo.New().NewContext(requestData, responseData)
		context.SetParamNames("userId", "collectionId")
		context.SetParamValues("1", "2")

		mockService := new(MockTaskService)
		taskHandler := Task{service: mockService}
		serviceErr := todoerrors.NewUnexpectedInternalError("Service layer error")
		mockService.On("FindByCollectionId", mock.Anything, mock.Anything).Return(nil,
			serviceErr)

		_ = taskHandler.FindByCollectionId(context)

		expectedBody := "{\"message\":\"Service layer error\"}\n"

		assert.Equal(t, http.StatusInternalServerError, responseData.Code)
		assert.Equal(t, expectedBody, responseData.Body.String())
	})
}
