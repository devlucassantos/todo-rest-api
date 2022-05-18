package postgres

import (
	"strings"
	"todo/src/common/crypto"
	"todo/src/core/domain"
	"todo/src/core/projecterrors/repositoryerrors"
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
	return &Auth{connectionManager}
}

func (r Auth) SignUp(account domain.Account) (int, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return -1, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	var id int
	err = connection.QueryRow(query.Auth().Insert(), dto.Account().Insert(account)...).Scan(&id)
	if err != nil {
		log.Error(err)
		return -1, r.handlePostgresError(err)
	}

	return id, nil
}

func (r Auth) SignIn(account domain.Account) (string, error) {
	conn, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return "", repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(conn)

	destination := dto.Account().Select().ByEmail()
	err = conn.Get(&destination, query.Auth().Select().ByEmail(), account.Email())
	if err != nil {
		return "", repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}
	isValid := crypto.ComparePassword(account.Password(), destination.Hash, destination.Password)
	if !isValid {
		return "", repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}
	token, err := destination.ConvertToDomain().GenerateToken()
	if err != nil {
		log.Error(err)
		return "", repositoryerrors.NewUnknownError(err)
	}
	stmt, err := conn.Prepare(query.Auth().UpdateToken())
	if err != nil {
		log.Error(err)
		return "", repositoryerrors.NewUnknownError(err)
	}

	_, err = stmt.Exec(dto.Account().UpdateToken(destination.Id, token)...)
	if err != nil {
		log.Error(err)
		return "", repositoryerrors.NewUnknownError(err)
	}

	return token, nil
}

func (r Auth) handlePostgresError(err error) error {
	errMessage := err.Error()

	if strings.Contains(errMessage, "unique") {
		return repositoryerrors.NewDuplicatedError(msgs.DuplicatedCredentials, err, msgs.Email)
	} else if strings.Contains(errMessage, "sql: no rows in result set") {
		return repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}

	return repositoryerrors.NewUnknownError(err)
}
