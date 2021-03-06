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

// Create
// @ID 			Create
// @Summary		Create a collection
// @Tags 		Collection
// @Description Route that allows registering a collection in the system. To register a collection it is necessary to inform the following data in the body of the request:
// @Description |   Name   |  Type  |   Required  | Description      |
// @Description |----------|--------|-------------|------------------|
// @Description |   name   | string |      x      | Collection name |
// @Accept 		json
// @Produce 	json
// @Security	bearerAuth
// @Param 	    userId       path       int                          true          "User ID"    default(1)
// @Param 		authJson 	 body 		request.Collection           true          "JSON responsible for sending all collection registration data to the database"
// @Success 	201 		 {object} 	response.SwaggerIdResponse                 "Collection successfully registered"
// @Failure 	400 		 {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401          {object}   response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	422 		 {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection  [post]
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

// Update
// @ID 			Update
// @Summary		Update a collection
// @Tags 		Collection
// @Description Route that allows editing a collection in the system. To edit a collection it is necessary to inform the following data:
// @Description |   Name   |  Type  |   Required  | Description	     |
// @Description |----------|--------|-------------|------------------|
// @Description |   name   | string |      x      | Collection name  |
// @Accept 		json
// @Produce 	json
// @Security	bearerAuth
// @Param 	    userId          path        int                          true          "User ID"          default(1)
// @Param 	    collectionId    path        int                          true          "Collection ID"    default(1)
// @Param 		authJson 	    body 	    request.Collection           true          "JSON responsible for sending the data needed to update the collection in the database"
// @Success 	204             {object}    nil 									   "Collection successfully edited"
// @Failure 	400             {object}    response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401             {object}    response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	404             {object}    response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422             {object}    response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500             {object}    response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection/{collectionId}  [put]
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

// Delete
// @ID 			Delete
// @Summary		Delete a collection
// @Tags 		Collection
// @Description Route that allows deleting a collection registered in the system
// @Security	bearerAuth
// @Param 	    userId          path    int                  true                  "User ID"          default(1)
// @Param 	    collectionId    path    int                  true                  "Collection ID"    default(1)
// @Success 	204 		 {object} 	nil                                        "Collection successfully deleted"
// @Failure 	400 		 {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401          {object}   response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	404 		 {object} 	response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422 		 {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection/{collectionId}  [delete]
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

// FindAll
// @ID 			FindAll
// @Summary 	Lists all user collections
// @Tags 		Collection
// @Description Route that allows searching all user collections in the system
// @Produce		json
// @Security	bearerAuth
// @Param 		userId    path      int                 true                   "User ID"    default(1)
// @Success 	200       {array} 	response.SwaggerCollectionResponse         "Successful request"
// @Failure 	400       {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401       {object} 	response.SwaggerUnauthorizedResponse       "The user is not authorized to make this request"
// @Failure 	404       {object} 	response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422       {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500       {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection 	[get]
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

// FindById
// @ID 			FindById
// @Summary 	Search a collection's data by ID
// @Tags 		Collection
// @Description Route that allows searching a collection registered in the system by ID
// @Produce		json
// @Security	bearerAuth
// @Param 	    userId          path        int                true                 "User ID"          default(1)
// @Param 	    collectionId    path        int                true                 "Collection ID"    default(1)
// @Success 	200          {object}    response.SwaggerCollectionResponse         "Successful request"
// @Failure 	400 		 {object} 	 response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401          {object}    response.SwaggerUnauthorizedResponse 	    "The user is not authorized to make this request"
// @Failure 	404 		 {object} 	 response.SwaggerNotFoundErrorResponse 	    "The user has requested a non-existent resource"
// @Failure 	422 		 {object} 	 response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	 response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection/{collectionId}    [get]
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
