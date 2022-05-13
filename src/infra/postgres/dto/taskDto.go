package dto

import "todo/src/core/domain"

type taskDto struct {
	Id             int    `db:"task_id"`
	Description    string `db:"task_description"`
	Finished       bool   `db:"task_finished"`
	CollectionId   int    `db:"collection_id"`
	CollectionName string `db:"collection_name"`
}

func (d taskDto) ConvertToDomain() *domain.Task {
	collection := domain.NewCollection(d.CollectionId, d.CollectionName)

	return domain.NewTask(d.Id, d.Description, d.Finished, collection)
}

type taskDtoManager struct{}

func Task() *taskDtoManager {
	return &taskDtoManager{}
}

func (taskDtoManager) Insert(task domain.Task) []interface{} {
	var collection *int
	collectionId := task.Collection().Id()
	if collectionId != 0 {
		collection = &collectionId
	}

	return []interface{}{
		task.Description(),
		task.Finished(),
		collection,
	}
}

func (taskDtoManager) Update(task domain.Task) []interface{} {
	var collection *int
	collectionId := task.Collection().Id()
	if collectionId != 0 {
		collection = &collectionId
	}

	return []interface{}{
		task.Description(),
		task.Finished(),
		collection,
		task.Id(),
	}
}

type taskDtoSelectManager struct{}

func (taskDtoManager) Select() *taskDtoSelectManager {
	return &taskDtoSelectManager{}
}

func (taskDtoSelectManager) All() []taskDto {
	return []taskDto{}
}

func (taskDtoSelectManager) ById() taskDto {
	return taskDto{}
}

func (taskDtoSelectManager) ByCollection() []taskDto {
	return []taskDto{}
}
