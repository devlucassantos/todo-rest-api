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

// Create
// @ID 			Create
// @Summary		Create a task
// @Tags 		Task
// @Description Route that allows registering a task in the system. To register a task it is necessary to inform the following data in the body of the request:
// @Description |      Name     |  Type  |   Required  |                    Description                    |
// @Description |---------------|--------|-------------|---------------------------------------------------|
// @Description | description   | string |             | Task description                                  |
// @Description | finished      |  bool  |             | If the task has been completed                    |
// @Description | collection_id |  int   |             | ID of the collection to which the task is related |
// @Accept 		json
// @Produce 	json
// @Security	bearerAuth
// @Param 	    userId       path       int                          true          "User ID"    default(1)
// @Param 		authJson 	 body 		request.Task                 true          "JSON responsible for sending all task registration data to the database"
// @Success 	201 		 {object} 	response.SwaggerIdResponse                 "Task successfully registered"
// @Failure 	400 		 {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401          {object}   response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	422 		 {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/task  [post]
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

// Update
// @ID 			Update
// @Summary		Update a task
// @Tags 		Task
// @Description Route that allows editing a task in the system. To edit a task it is necessary to inform the following data:
// @Description |      Name     |  Type  |   Required  |                    Description                    |
// @Description |---------------|--------|-------------|---------------------------------------------------|
// @Description | description   | string |             | Task description                                  |
// @Description | finished      |  bool  |             | If the task has been completed                    |
// @Description | collection_id |  int   |             | ID of the collection to which the task is related |
// @Accept 		json
// @Produce 	json
// @Security	bearerAuth
// @Param 	    userId      path        int                          true          "User ID"          default(1)
// @Param 	    taskId      path        int                          true          "Task ID"    default(1)
// @Param 		authJson    body 	    request.Task                 true          "JSON responsible for sending the data needed to update the task in the database"
// @Success 	204         {object}    nil 									   "Task successfully edited"
// @Failure 	400         {object}    response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401         {object}    response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	404         {object}    response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422         {object}    response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500         {object}    response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/task/{taskId}  [put]
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

// Delete
// @ID 			Delete
// @Summary		Delete a task
// @Tags 		Task
// @Description Route that allows deleting a task registered in the system
// @Security	bearerAuth
// @Param 	    userId       path       int                  true                  "User ID"    default(1)
// @Param 	    taskId       path       int                  true                  "Task ID"    default(1)
// @Success 	204 		 {object} 	nil                                        "Task successfully deleted"
// @Failure 	400 		 {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401          {object}   response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	404 		 {object} 	response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422 		 {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/task/{taskId}  [delete]
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

// FindAll
// @ID 			FindAll
// @Summary 	Lists all user tasks
// @Tags 		Task
// @Description Route that allows searching all user tasks in the system
// @Produce		json
// @Security	bearerAuth
// @Param 		userId    path      int                 true                   "User ID"    default(1)
// @Success 	200       {array} 	response.SwaggerTaskResponse               "Successful request"
// @Failure 	400       {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401       {object} 	response.SwaggerUnauthorizedResponse       "The user is not authorized to make this request"
// @Failure 	404       {object} 	response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422       {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500       {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/task 	[get]
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

// FindById
// @ID 			FindById
// @Summary 	Search a task's data by ID
// @Tags 		Task
// @Description Route that allows searching a task registered in the system by ID
// @Produce		json
// @Security	bearerAuth
// @Param 	    userId    path        int                true                    "User ID"    default(1)
// @Param 	    taskId    path        int                true                    "Task ID"    default(1)
// @Success 	200       {object}    response.SwaggerTaskResponse               "Successful request"
// @Failure 	400       {object}    response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401       {object}    response.SwaggerUnauthorizedResponse 	     "The user is not authorized to make this request"
// @Failure 	404       {object} 	  response.SwaggerNotFoundErrorResponse 	 "The user has requested a non-existent resource"
// @Failure 	422       {object} 	  response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500       {object} 	  response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/task/{taskId}    [get]
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

// FindByCollectionId
// @ID 			FindByCollectionId
// @Summary 	Search all tasks by collection ID
// @Tags 		Collection
// @Description Route that allows searching all tasks registered in the system by collection ID
// @Produce		json
// @Security	bearerAuth
// @Param 	    userId          path        int                true                    "User ID"          default(1)
// @Param 	    collectionId    path        int                true                    "Collection ID"    default(1)
// @Success 	200             {object}    response.SwaggerTaskResponse               "Successful request"
// @Failure 	400             {object}    response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	401             {object}    response.SwaggerUnauthorizedResponse 	   "The user is not authorized to make this request"
// @Failure 	404             {object} 	response.SwaggerNotFoundErrorResponse 	   "The user has requested a non-existent resource"
// @Failure 	422             {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500             {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/user/{userId}/collection/{collectionId}/tasks    [get]
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
