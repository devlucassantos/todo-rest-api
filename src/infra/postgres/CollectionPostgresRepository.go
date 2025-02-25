package postgres

import (
	"errors"
	"github.com/labstack/gommon/log"
	"strings"
	"todo/src/core/domain"
	"todo/src/core/projecterrors/repositoryerrors"
	"todo/src/infra/postgres/dto"
	"todo/src/infra/postgres/msgs"
	"todo/src/infra/postgres/query"
)

type Collection struct {
	iConnectionManager
}

func NewCollectionPostgresRepository(connectionManager iConnectionManager) *Collection {
	return &Collection{
		connectionManager,
	}
}

func (r Collection) Create(collection domain.Collection, userId int) (int, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return -1, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	var id int
	err = connection.QueryRow(query.Collection().Insert(), dto.Collection().Insert(collection, userId)...).Scan(&id)
	if err != nil {
		log.Error(err)
		return -1, r.handlePostgresError(err)
	}

	return id, nil
}

func (r Collection) Update(collection domain.Collection, userId int) error {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	result, err := connection.Exec(query.Collection().Update(), dto.Collection().Update(collection, userId)...)
	if err != nil {
		log.Error(err)
		return r.handlePostgresError(err)
	}
	affectedRows, resultErr := result.RowsAffected()
	if affectedRows == 0 {
		return repositoryerrors.NewNotFoundError(msgs.CollectionNotFound, errors.New(msgs.CollectionNotFoundNewError))
	} else if resultErr != nil {
		log.Error(resultErr)
		return repositoryerrors.NewUnknownError(resultErr)
	}

	return nil
}

func (r Collection) Delete(collectionId, userId int) error {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	result, err := connection.Exec(query.Collection().Delete(), collectionId, userId)
	if err != nil {
		log.Error(err)
		return r.handlePostgresError(err)
	}
	if affectedRows, resultErr := result.RowsAffected(); affectedRows == 0 {
		return repositoryerrors.NewNotFoundError(msgs.CollectionNotFound, errors.New(msgs.CollectionNotFoundNewError))
	} else if resultErr != nil {
		log.Error(resultErr)
		return repositoryerrors.NewUnknownError(resultErr)
	}

	return nil
}

func (r Collection) FindAll(userId int) ([]domain.Collection, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	destination := dto.Collection().Select().All()
	err = connection.Select(&destination, query.Collection().Select().All(), userId)
	if err != nil {
		log.Error(err)
		return nil, r.handlePostgresError(err)
	}
	var collectionList []domain.Collection
	for _, collection := range destination {
		collectionList = append(collectionList, *collection.ConvertToDomain())
	}

	return collectionList, nil
}

func (r Collection) handlePostgresError(err error) error {
	errMessage := err.Error()

	if strings.Contains(errMessage, "sql: no rows in result set") {
		return repositoryerrors.NewNotFoundError(msgs.CollectionNotFound, err)
	}

	return repositoryerrors.NewUnknownError(err)
}
