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

type Collection struct {
	service interfaces.ICollection
}

func NewCollectionHandler() *Collection {
	connectionManager := postgres.NewPostgresConnectionManager()
	repository := postgres.NewCollectionPostgresRepository(connectionManager)
	service := services.NewCollectionService(repository)
	return &Collection{service}
}

func (h Collection) Create(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	var requestData request.Collection
	if err = ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	collection, collectionErr := domain.NewValidatedCollection(
		-1,
		requestData.Name,
	)
	if collectionErr != nil {
		log.Error(collectionErr)
		return writeValidationError(ctx, *todoerrors.NewValidationError(collectionErr.Error(), *collectionErr.InvalidFields()))
	}

	userIdCreated, err := h.service.Create(*collection, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	responseReturned := map[string]int{"id": userIdCreated}
	return writeCreatedResponse(ctx, responseReturned)
}

func (h Collection) Update(ctx echo.Context) error {
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
	var requestData request.Collection
	if err = ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	collection, collectionErr := domain.NewValidatedCollection(
		collectionId,
		requestData.Name,
	)
	if collectionErr != nil {
		log.Error(collectionErr)
		return writeValidationError(ctx, *todoerrors.NewValidationError(collectionErr.Error(), *collectionErr.InvalidFields()))
	}

	err = h.service.Update(*collection, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	return writeNoContentResponse(ctx)
}

func (h Collection) Delete(ctx echo.Context) error {
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

	err = h.service.Delete(collectionId, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	return writeNoContentResponse(ctx)
}

func (h Collection) FindAll(ctx echo.Context) error {
	userId, err := convertToInt(ctx.Param("userId"), msgs.UserId)
	if err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.UserId, msgs.ConversionError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}

	collectionList, err := h.service.FindAll(userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	var collectionResponseList []response.Collection
	for _, collection := range collectionList {
		collectionResponseList = append(collectionResponseList, *response.NewCollection(collection))
	}
	return writeAcceptResponse(ctx, collectionResponseList)
}

func (h Collection) FindById(ctx echo.Context) error {
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

	collection, err := h.service.FindById(collectionId, userId)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	collectionResponse := response.NewCollection(*collection)
	return writeAcceptResponse(ctx, collectionResponse)
}
