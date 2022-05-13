package repository

import "todo/src/core/domain"

type IAuth interface {
	SignUp(account domain.Account) (int, error)
	SignIn(account domain.Account) (string, error)
}
