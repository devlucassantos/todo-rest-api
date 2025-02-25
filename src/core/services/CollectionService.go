package services

import (
	"github.com/labstack/gommon/log"
	"todo/src/core/domain"
	"todo/src/core/interfaces/repository"
	"todo/src/core/projecterrors/todoerrors"
)

type Collection struct {
	repository repository.ICollection
}

func NewCollectionService(repository repository.ICollection) *Collection {
	return &Collection{repository}
}

func (s Collection) Create(collection domain.Collection, userId int) (int, error) {
	id, err := s.repository.Create(collection, userId)
	if err != nil {
		log.Error(err)
		return -1, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Create)
	}

	return id, nil
}

func (s Collection) Update(collection domain.Collection, userId int) error {
	err := s.repository.Update(collection, userId)
	if err != nil {
		log.Error(err)
		return todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Update)
	}

	return nil
}

func (s Collection) Delete(collectionId, userId int) error {
	err := s.repository.Delete(collectionId, userId)
	if err != nil {
		log.Error(err)
		return todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.Delete)
	}

	return nil
}

func (s Collection) FindAll(userId int) ([]domain.Collection, error) {
	collectionList, err := s.repository.FindAll(userId)
	if err != nil {
		log.Error(err)
		return nil, todoerrors.ConvertRepositoryErrorToServiceError(err, s.repository.FindAll)
	}

	return collectionList, nil
}
