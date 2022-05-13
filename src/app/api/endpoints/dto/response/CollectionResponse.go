package response

import "todo/src/core/domain"

type Collection struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewCollection(collection domain.Collection) *Collection {
	return &Collection{
		Id:   collection.Id(),
		Name: collection.Name(),
	}
}
