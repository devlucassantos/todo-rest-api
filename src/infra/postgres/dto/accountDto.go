package dto

import (
	"todo/src/core/domain"
)

type accountDto struct {
	Id       int    `db:"account_id"`
	Name     string `db:"account_name"`
	Email    string `db:"account_email"`
	Password string `db:"account_password"`
	Token    string `db:"account_token"`
}

func (d accountDto) ConvertToDomain() *domain.Account {
	return domain.NewAccount(
		d.Id,
		d.Name,
		d.Email,
		d.Password,
		d.Token,
	)
}

type accountDtoManager struct{}

func Account() *accountDtoManager {
	return &accountDtoManager{}
}

func (accountDtoManager) Insert(account domain.Account) []interface{} {
	return []interface{}{
		account.Name(),
		account.Email(),
		account.Password(),
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
