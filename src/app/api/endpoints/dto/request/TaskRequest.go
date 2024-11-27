package request

type Task struct {
	Description  string `json:"description"`
	Finished     bool   `json:"finished"`
	CollectionId int    `json:"collection_id"`
}
