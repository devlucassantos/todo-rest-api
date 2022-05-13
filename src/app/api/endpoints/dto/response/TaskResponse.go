package response

import "todo/src/core/domain"

type Task struct {
	Id          int         `json:"id"`
	Description string      `json:"description"`
	Finished    bool        `json:"finished"`
	Collection  *Collection `json:"collection"`
}

func NewTask(task domain.Task) *Task {
	collection := NewCollection(*task.Collection())

	return &Task{
		Id:          task.Id(),
		Description: task.Description(),
		Finished:    task.Finished(),
		Collection:  collection,
	}
}
