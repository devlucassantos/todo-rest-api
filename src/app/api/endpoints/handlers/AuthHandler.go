package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"todo/src/app/api/endpoints/dto/request"
	"todo/src/app/api/endpoints/dto/response"
	"todo/src/core/domain"
	"todo/src/core/errs/handlererrs"
	iServices "todo/src/core/interfaces/services"
	"todo/src/core/services"
	"todo/src/infra/postgres"
)

type Auth struct {
	service iServices.IAuth
}

func NewAuthHandler() *Auth {
	connectionManager := postgres.NewPostgresConnectionManager()
	repository := postgres.NewAuthPostgresRepository(connectionManager)
	service := services.NewAuthService(repository)
	return &Auth{service}
}

func (h Auth) SignUp(ctx echo.Context) error {
	var accountRequest request.Account
	err := ctx.Bind(&accountRequest)
	if err != nil {
		log.Error(err)
		errMessage, invalidFields := handlererrs.GetAuthBindError(err)
		return writeValidationErr(ctx, handlererrs.NewValidationErr(errMessage, invalidFields))
	}

	account := domain.NewAccount(
		-1,
		accountRequest.Name,
		accountRequest.Email,
		accountRequest.Password,
		"",
		"",
	)

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
	var accountRequest request.Account
	err := ctx.Bind(&accountRequest)
	if err != nil {
		log.Error(err)
		errMessage, invalidFields := handlererrs.GetAuthBindError(err)
		return writeValidationErr(ctx, handlererrs.NewValidationErr(errMessage, invalidFields))
	}

	account := domain.NewAccount(
		-1,
		"",
		accountRequest.Email,
		accountRequest.Password,
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
