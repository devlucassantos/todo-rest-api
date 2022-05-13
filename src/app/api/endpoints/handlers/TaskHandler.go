package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"todo/src/app/api/endpoints/dto/request"
	"todo/src/app/api/endpoints/dto/response"
	"todo/src/app/api/endpoints/handlers/msgs"
	"todo/src/core/domain"
	interfaces "todo/src/core/interfaces/services"
	"todo/src/core/projecterrors/todoerrors"
	"todo/src/core/services"
	"todo/src/infra/postgres"
)

type Task struct {
	service interfaces.ITask
}

func NewTaskHandler() *Task {
	connectionManager := postgres.NewPostgresConnectionManager()
	repository := postgres.NewTaskPostgresRepository(connectionManager)
	service := services.NewTaskService(repository)
	return &Task{service}
}

func (h Task) Create(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	var requestData request.Task
	if err = ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	collection := domain.NewCollection(requestData.CollectionId, "")
	task := domain.NewTask(
		-1,
		requestData.Description,
		requestData.Finished,
		collection,
	)

	userIdCreated, err := h.service.Create(*task, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	responseReturned := map[string]int{"id": userIdCreated}
	return writeCreatedResponse(ctx, responseReturned)
}

func (h Task) Update(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	taskId, err := convertToInt(ctx.Param("taskId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.TaskId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	var requestData request.Task
	if err = ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	collection := domain.NewCollection(requestData.CollectionId, "")
	task := domain.NewTask(
		taskId,
		requestData.Description,
		requestData.Finished,
		collection,
	)

	err = h.service.Update(*task, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	return writeNoContentResponse(ctx)
}

func (h Task) Delete(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	taskId, err := convertToInt(ctx.Param("taskId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.TaskId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}

	err = h.service.Delete(taskId, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	return writeNoContentResponse(ctx)
}

func (h Task) FindAll(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}

	taskList, err := h.service.FindAll(userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	var taskResponseList []response.Task
	for _, task := range taskList {
		taskResponseList = append(taskResponseList, *response.NewTask(task))
	}
	return writeAcceptResponse(ctx, taskResponseList)
}

func (h Task) FindById(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	taskId, err := convertToInt(ctx.Param("taskId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.TaskId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}

	task, err := h.service.FindById(taskId, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	taskResponse := response.NewTask(*task)
	return writeAcceptResponse(ctx, taskResponse)
}

func (h Task) FindByCollectionId(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	collectionId, err := convertToInt(ctx.Param("collectionId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.CollectionId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}

	taskList, err := h.service.FindByCollectionId(collectionId, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	var taskResponseList []response.Task
	for _, task := range taskList {
		taskResponseList = append(taskResponseList, *response.NewTask(task))
	}
	return writeAcceptResponse(ctx, taskResponseList)
}
