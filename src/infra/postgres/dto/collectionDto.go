package dto

import "todo/src/core/domain"

type collectionDto struct {
	Id   int    `db:"collection_id"`
	Name string `db:"collection_name"`
}

func (d collectionDto) ConvertToDomain() *domain.Collection {
	return domain.NewCollection(d.Id, d.Name)
}

type collectionDtoManager struct{}

func Collection() *collectionDtoManager {
	return &collectionDtoManager{}
}

func (collectionDtoManager) Insert(collection domain.Collection, userId int) []interface{} {
	return []interface{}{
		collection.Name(),
		userId,
	}
}

func (collectionDtoManager) Update(collection domain.Collection, userId int) []interface{} {
	return []interface{}{
		collection.Name(),
		collection.Id(),
		userId,
	}
}

type collectionDtoSelectManager struct{}

func (collectionDtoManager) Select() *collectionDtoSelectManager {
	return &collectionDtoSelectManager{}
}

func (collectionDtoSelectManager) All() []collectionDto {
	return []collectionDto{}
}

func (collectionDtoSelectManager) ById() collectionDto {
	return collectionDto{}
}
