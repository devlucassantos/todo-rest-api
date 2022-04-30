package repository

import "todo/src/core/domain"

type IAuth interface {
	SignUp(domain.Account) (int, error)
	SignIn(domain.Account) (string, error)
}
