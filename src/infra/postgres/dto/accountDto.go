package dto

import (
	"todo/src/core/domain"
)

type accountDto struct {
	Id       int    `db:"account_id"`
	Name     string `db:"account_name"`
	Email    string `db:"account_email"`
	Password string `db:"account_password"`
	Hash     string `db:"account_hash"`
	Token    string `db:"account_token"`
}

func (a accountDto) ConvertToDomain() *domain.Account {
	return domain.NewAccount(
		a.Id,
		a.Name,
		a.Email,
		a.Password,
		a.Hash,
		a.Token,
	)
}

func Account() *accountDtoManager {
	return &accountDtoManager{}
}

type accountDtoManager struct{}

func (accountDtoManager) Insert(account domain.Account) []interface{} {
	return []interface{}{
		account.Name(),
		account.Email(),
		account.Password(),
		account.Hash(),
		account.Token(),
	}
}

func (accountDtoManager) UpdateToken(id int, token string) []interface{} {
	return []interface{}{
		token,
		id,
	}
}

type accountDtoSelectManager struct{}

func (accountDtoManager) Select() *accountDtoSelectManager {
	return &accountDtoSelectManager{}
}

func (accountDtoSelectManager) ByEmail() accountDto {
	return accountDto{}
}
