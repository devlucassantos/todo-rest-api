package postgres

import (
	"strings"
	"todo/src/common/crypto"
	"todo/src/core/domain"
	"todo/src/core/errs/repositoryerrs"
	"todo/src/infra/postgres/dto"
	"todo/src/infra/postgres/msgs"
	"todo/src/infra/postgres/query"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Auth struct {
	iConnectionManager
}

func NewAuthPostgresRepository(connectionManager iConnectionManager) *Auth {
	return &Auth{
		connectionManager,
	}
}

func (r Auth) SignUp(account domain.Account) (int, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return -1, repositoryerrs.NewServiceUnavailableErr(msgs.ConnectionErr, err)
	}
	defer r.closeConnection(connection)

	var id int
	err = connection.QueryRow(query.Auth().Insert(), dto.Account().Insert(account)...).Scan(&id)
	if err != nil {
		log.Error(err)
		return -1, handlePostgresError(err)
	}

	return id, nil
}

func (r Auth) SignIn(account domain.Account) (string, error) {
	conn, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return "", repositoryerrs.NewServiceUnavailableErr(msgs.ConnectionErr, err)
	}
	defer r.closeConnection(conn)

	destination := dto.Account().Select().ByEmail()
	err = conn.Get(&destination, query.Auth().ByEmail(), account.Email())
	if err != nil {
		return "", repositoryerrs.NewNotFoundErr(msgs.InvalidAccountCredentials, err)
	}

	isValid := crypto.ComparePassword(account.Password(), destination.Hash, destination.Password)
	if !isValid {
		return "", repositoryerrs.NewNotFoundErr(msgs.InvalidAccountCredentials, err)
	}

	token, err := account.GenerateToken()
	if err != nil {
		log.Error(err)
		return "", repositoryerrs.NewUnknownErr(err)
	}

	stmt, err := conn.Prepare(query.Auth().UpdateToken())
	if err != nil {
		log.Error(err)
		return "", repositoryerrs.NewUnknownErr(err)
	}

	_, err = stmt.Exec(dto.Account().UpdateToken(destination.Id, token)...)
	if err != nil {
		log.Error(err)
		return "", repositoryerrs.NewUnknownErr(err)
	}

	return token, nil
}

func handlePostgresError(err error) error {
	errMessage := err.Error()

	if strings.Contains(errMessage, "unique") {
		return repositoryerrs.NewDuplicatedErr(msgs.DuplicatedCredentials, err)
	} else if strings.Contains(errMessage, "sql: no rows in result set") {
		return repositoryerrs.NewNotFoundErr(msgs.InvalidAccountCredentials, err)
	}

	return repositoryerrs.NewUnknownErr(err)
}
