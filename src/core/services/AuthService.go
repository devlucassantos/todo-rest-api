package services

import (
	"encoding/hex"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
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

func (s Auth) SignUp(account domain.Account) (*domain.Account, error) {
	if account.Name() == "" || account.Email() == "" || account.Password() == "" {
		return nil, todoerrors.NewMissingInfoError(msgs.EmptyNameEmailOrPassword)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password()), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: " + err.Error())
		return nil, err
	}

	account.SetPassword(hex.EncodeToString(hashedPassword))

	accountData, err := s.repository.SignUp(account)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.SignUp)
	}

	token, err := account.GenerateToken()
	if err != nil {
		log.Error(err)
		return nil, todoerrors.NewUnexpectedInternalError(errmsgs.UnexpectedInternalError)
	}

	account = *domain.NewAccount(
		accountData.Id(),
		accountData.Name(),
		accountData.Email(),
		"",
		token,
	)

	return &account, nil
}

func (s Auth) SignIn(account domain.Account) (*domain.Account, error) {
	if account.Email() == "" || account.Password() == "" {
		return nil, todoerrors.NewMissingInfoError(msgs.EmptyEmailOrPassword)
	}

	accountData, err := s.repository.SignIn(account)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.SignIn)
	}

	token, err := accountData.GenerateToken()
	if err != nil {
		log.Error(err)
		return nil, todoerrors.NewUnexpectedInternalError(errmsgs.UnexpectedInternalError)
	}

	account = *domain.NewAccount(
		accountData.Id(),
		accountData.Name(),
		accountData.Email(),
		"",
		token,
	)

	return &account, nil
}
