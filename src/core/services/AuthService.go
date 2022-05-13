package services

import (
	"github.com/labstack/gommon/log"
	"todo/src/common/crypto"
	"todo/src/core/domain"
	"todo/src/core/interfaces/repository"
	"todo/src/core/projecterrors/todoerrors"
	errmsgs "todo/src/core/projecterrors/todoerrors/msgs"
	"todo/src/core/services/msgs"
)

type Auth struct {
	repository repository.IAuth
}

func NewAuthService(repository repository.IAuth) *Auth {
	return &Auth{repository}
}

func (s Auth) SignUp(account domain.Account) (*int, *string, error) {
	if account.Name() == "" || account.Email() == "" || account.Password() == "" {
		return nil, nil, todoerrors.NewMissingInfoError(msgs.EmptyNameEmailOrPassword)
	}

	password, hash, err := crypto.HashPassword(account.Password())
	if err != nil {
		log.Error(err)
		return nil, nil, todoerrors.NewUnexpectedInternalError(errmsgs.UnexpectedInternalError)
	}

	account.SetPassword(password)
	account.SetHash(hash)

	id, err := s.repository.SignUp(account)
	if err != nil {
		log.Error(err)
		return nil, nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.SignUp)
	}

	token, err := account.GenerateToken()
	if err != nil {
		log.Error(err)
		return nil, nil, todoerrors.NewUnexpectedInternalError(errmsgs.UnexpectedInternalError)
	}

	return &id, &token, nil
}

func (s Auth) SignIn(account domain.Account) (*string, error) {
	if account.Email() == "" || account.Password() == "" {
		return nil, todoerrors.NewMissingInfoError(msgs.EmptyEmailOrPassword)
	}

	token, err := s.repository.SignIn(account)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.SignIn)
	}

	return &token, nil
}
