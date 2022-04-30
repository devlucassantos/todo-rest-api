package services

import (
	"github.com/labstack/gommon/log"
	"todo/src/common/crypto"
	"todo/src/core/domain"
	"todo/src/core/errs/serviceerrs"
	"todo/src/core/errs/serviceerrs/msgs"
	"todo/src/core/interfaces/repository"
)

type Auth struct {
	repository repository.IAuth
}

func NewAuthService(repository repository.IAuth) *Auth {
	return &Auth{repository}
}

func (s Auth) SignUp(account domain.Account) (*int, *string, error) {
	if account.Name() == "" || account.Email() == "" || account.Password() == "" {
		return nil, nil, serviceerrs.NewMissingInfoErr("The user name, email and password must not be empty.")
	}

	password, hash, err := crypto.HashPassword(account.Password())
	if err != nil {
		log.Error(err)
		return nil, nil, serviceerrs.NewUnexpectedInternalErr(msgs.UnexpectedInternalErr)
	}

	account.SetPassword(password)
	account.SetHash(hash)

	id, err := s.repository.SignUp(account)
	if err != nil {
		log.Error(err)
		return nil, nil, serviceerrs.ConvertRepositoryErrToServiceErr(err, s.SignUp)
	}

	token, err := account.GenerateToken()
	if err != nil {
		log.Error(err)
		return nil, nil, serviceerrs.NewUnexpectedInternalErr(msgs.UnexpectedInternalErr)
	}

	return &id, &token, nil
}

func (s Auth) SignIn(account domain.Account) (*string, error) {
	if account.Email() == "" || account.Password() == "" {
		return nil, serviceerrs.NewMissingInfoErr("The email and password must not be empty.")
	}

	token, err := s.repository.SignIn(account)
	if err != nil {
		log.Error(err)
		return nil, serviceerrs.ConvertRepositoryErrToServiceErr(err, s.SignIn)
	}

	return &token, nil
}
