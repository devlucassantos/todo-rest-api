package services

import "todo/src/core/domain"

type ITask interface {
	Create(task domain.Task, userId int) (int, error)
	Update(task domain.Task, userId int) error
	Delete(taskId, userId int) error
	FindAll(userId int) ([]domain.Task, error)
	FindByCollectionId(collectionId, userId int) ([]domain.Task, error)
}
