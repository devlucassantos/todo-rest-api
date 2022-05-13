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
