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

type Auth struct {
	service interfaces.IAuth
}

func NewAuthHandler() *Auth {
	connectionManager := postgres.NewPostgresConnectionManager()
	repository := postgres.NewAuthPostgresRepository(connectionManager)
	service := services.NewAuthService(repository)
	return &Auth{service}
}

// SignUp
// @ID 			SignUp
// @Summary		User Sign Up
// @Tags 		Authentication
// @Description Route that allows you to register a user in the system. To register a user it is necessary to inform the following data in the body of the request:
// @Description |   Name   |  Type  |   Required  | Description		|
// @Description |----------|--------|-------------|-----------------|
// @Description | name     | string |      x      | Real user name  |
// @Description | email    | string |      x      | User email      |
// @Description | password | string |      x      | User password   |
// @Accept 		json
// @Produce 	json
// @Param 		authJson 		body 		request.Account            true            "JSON responsible for sending all user registration data to the server"
// @Success 	201 			{object} 	response.SwaggerSignUpResponse 			   "User successfully registered"
// @Failure 	400 			{object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	409 			{object} 	response.SwaggerConflictErrorResponse 	   "The user tried to register with the email of an existing user"
// @Failure 	422 			{object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 			{object} 	response.SwaggerGenericErrorResponse 	   "An unexpected server error has occurred"
// @Router 		/auth/signup 	[post]
func (h Auth) SignUp(ctx echo.Context) error {
	var requestData request.Account
	if err := ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	account, validationError := domain.NewValidatedAccount(
		-1,
		requestData.Name,
		requestData.Email,
		requestData.Password,
		"",
		"",
	)
	if validationError != nil {
		return writeValidationError(ctx, *validationError)
	}

	id, token, err := h.service.SignUp(*account)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	authResponse := response.Auth{
		Id:          *id,
		AccessToken: *token,
	}
	return writeCreatedResponse(ctx, authResponse)
}

// SignIn
// @ID 			SignIn
// @Summary		User Sign In
// @Tags 		Authentication
// @Description Route that allows connecting the user to the system through their registration data. To connect a user it is necessary to inform the following data in the body of the request:
// @Description |   Name   |  Type  |   Required  | Description		|
// @Description |----------|--------|-------------|-----------------|
// @Description | email    | string |      x      | User email      |
// @Description | password | string |      x      | User password   |
// @Accept 		json
// @Produce 	json
// @Param 		authJson 	 body 		request.Account            true            "JSON responsible for sending all user sign in data to the server"
// @Success 	201 		 {object} 	response.SwaggerLoginResponse 			   "User successfully signed in"
// @Failure 	400 		 {object} 	response.SwaggerValidationErrorResponse    "The user has made a bad request"
// @Failure 	422 		 {object} 	response.SwaggerValidationErrorResponse    "Some entered data could not be processed because it is not valid"
// @Failure 	500 		 {object} 	response.SwaggerGenericErrorResponse       "An unexpected server error has occurred"
// @Router 		/auth/signin [post]
func (h Auth) SignIn(ctx echo.Context) error {
	var requestData request.Account
	if err := ctx.Bind(&requestData); err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.Request, msgs.RequestFormatError)
		return writeValidationError(ctx, *todoerrors.NewValidationError(err.Error(), invalidFields))
	}
	account := domain.NewAccount(
		-1,
		"",
		requestData.Email,
		requestData.Password,
		"",
		"",
	)

	token, err := h.service.SignIn(*account)
	if err != nil {
		log.Error(err)
		return handleServiceErrors(ctx, err)
	}

	authResponse := response.Auth{
		AccessToken: *token,
	}
	return writeCreatedResponse(ctx, authResponse)
}
