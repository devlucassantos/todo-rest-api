package repository

import "todo/src/core/domain"

type IAuth interface {
	SignUp(account domain.Account) (*domain.Account, error)
	SignIn(account domain.Account) (*domain.Account, error)
}
