package services

import "todo/src/core/domain"

type ICollection interface {
	Create(collection domain.Collection, userId int) (int, error)
	Update(collection domain.Collection, userId int) error
	Delete(collectionId, userId int) error
	FindAll(userId int) ([]domain.Collection, error)
	FindById(collectionId, userId int) (*domain.Collection, error)
}
