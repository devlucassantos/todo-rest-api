package services

import (
	"github.com/labstack/gommon/log"
	"todo/src/core/domain"
	"todo/src/core/interfaces/repository"
	"todo/src/core/projecterrors/todoerrors"
)

type Task struct {
	repository repository.ITask
}

func NewTaskService(repository repository.ITask) *Task {
	return &Task{repository}
}

func (s Task) Create(task domain.Task, userId int) (int, error) {
	id, err := s.repository.Create(task, userId)
	if err != nil {
		log.Error(err)
		return -1, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Create)
	}

	return id, nil
}

func (s Task) Update(task domain.Task, userId int) error {
	err := s.repository.Update(task, userId)
	if err != nil {
		log.Error(err)
		return todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Update)
	}

	return nil
}

func (s Task) Delete(taskId, userId int) error {
	err := s.repository.Delete(taskId, userId)
	if err != nil {
		log.Error(err)
		return todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Delete)
	}

	return nil
}

func (s Task) FindAll(userId int) ([]domain.Task, error) {
	taskList, err := s.repository.FindAll(userId)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.FindAll)
	}

	return taskList, nil
}

func (s Task) FindById(taskId, userId int) (*domain.Task, error) {
	task, err := s.repository.FindById(taskId, userId)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.FindById)
	}

	return task, nil
}

func (s Task) FindByCollectionId(collectionId, userId int) ([]domain.Task, error) {
	taskList, err := s.repository.FindByCollectionId(collectionId, userId)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.FindByCollectionId)
	}

	return taskList, nil
}
