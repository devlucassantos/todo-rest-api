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

type Task struct {
	iConnectionManager
}

func NewTaskPostgresRepository(connectionManager iConnectionManager) *Task {
	return &Task{
		connectionManager,
	}
}

func (r Task) Create(task domain.Task, userId int) (int, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return -1, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	var id int
	insertData := append(dto.Task().Insert(task), userId)
	err = connection.QueryRow(query.Task().Insert(), insertData...).Scan(&id)
	if err != nil {
		log.Error(err)
		return -1, r.handlePostgresError(err)
	}

	return id, nil
}

func (r Task) Update(task domain.Task, userId int) error {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	insertData := append(dto.Task().Update(task), userId)
	result, err := connection.Exec(query.Task().Update(), insertData...)
	if err != nil {
		log.Error(err)
		return r.handlePostgresError(err)
	}
	affectedRows, resultErr := result.RowsAffected()
	if affectedRows == 0 {
		return repositoryerrors.NewNotFoundError(msgs.TaskNotFound, errors.New(msgs.TaskNotFoundNewError))
	} else if resultErr != nil {
		log.Error(resultErr)
		return repositoryerrors.NewUnknownError(resultErr)
	}

	return nil
}

func (r Task) Delete(taskId, userId int) error {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	result, err := connection.Exec(query.Task().Delete(), taskId, userId)
	if err != nil {
		log.Error(err)
		return r.handlePostgresError(err)
	}
	if affectedRows, resultErr := result.RowsAffected(); affectedRows == 0 {
		return repositoryerrors.NewNotFoundError(msgs.TaskNotFound, errors.New(msgs.TaskNotFoundNewError))
	} else if resultErr != nil {
		log.Error(resultErr)
		return repositoryerrors.NewUnknownError(resultErr)
	}

	return nil
}

func (r Task) FindAll(userId int) ([]domain.Task, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	destination := dto.Task().Select().All()
	err = connection.Select(&destination, query.Task().Select().All(), userId)
	if err != nil {
		log.Error(err)
		return nil, r.handlePostgresError(err)
	}
	var taskList []domain.Task
	for _, task := range destination {
		taskList = append(taskList, *task.ConvertToDomain())
	}

	return taskList, nil
}

func (r Task) FindById(taskId, userId int) (*domain.Task, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	destination := dto.Task().Select().ById()
	err = connection.Get(&destination, query.Task().Select().ById(), taskId, userId)
	if err != nil {
		log.Error(err)
		return nil, r.handlePostgresError(err)
	}

	return destination.ConvertToDomain(), nil
}

func (r Task) FindByCollectionId(collectionId, userId int) ([]domain.Task, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	destination := dto.Task().Select().ByCollection()
	err = connection.Select(&destination, query.Task().Select().ByCollection(), collectionId, userId)
	if err != nil {
		log.Error(err)
		return nil, r.handlePostgresError(err)
	}
	var taskList []domain.Task
	for _, task := range destination {
		taskList = append(taskList, *task.ConvertToDomain())
	}

	return taskList, nil
}

func (r Task) handlePostgresError(err error) error {
	errMessage := err.Error()

	if strings.Contains(errMessage, "unique") {
		return repositoryerrors.NewDuplicatedError(msgs.DuplicatedCredentials, err)
	} else if strings.Contains(errMessage, "task_collection_fk") {
		return repositoryerrors.NewDependencyError(msgs.CollectionNotFound, err, msgs.Collection)
	} else if strings.Contains(errMessage, "sql: no rows in result set") {
		return repositoryerrors.NewNotFoundError(msgs.InvalidAccountCredentials, err)
	}

	return repositoryerrors.NewUnknownError(err)
}
