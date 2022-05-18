package request

type Task struct {
	Description  string `json:"description"   example:"Task example"`
	Finished     bool   `json:"finished"      example:"false"`
	CollectionId int    `json:"collection_id" example:"1"`
}
