package domain

type Task struct {
	id          int
	description string
	finished    bool
	collection  *Collection
}

func NewTask(id int, description string, finished bool, collection *Collection) *Task {
	return &Task{
		id:          id,
		description: description,
		finished:    finished,
		collection:  collection,
	}
}

func (d Task) Id() int {
	return d.id
}

func (d Task) Description() string {
	return d.description
}

func (d Task) Finished() bool {
	return d.finished
}

func (d Task) Collection() *Collection {
	return d.collection
}
