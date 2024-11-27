package request

type SwaggerSignUpRequest struct {
	Name     string `json:"name"     example:"Example Name"`
	Email    string `json:"email"    example:"example@example.com"`
	Password string `json:"password" example:"ex@mplePassw0rd"`
}

type SwaggerSignInRequest struct {
	Email    string `json:"email"    example:"example@example.com"`
	Password string `json:"password" example:"ex@mplePassw0rd"`
}

type SwaggerCollectionRequest struct {
	Name string `json:"name" example:"Collection example"`
}

type SwaggerTaskRequest struct {
	Description  string `json:"description"   example:"Task example"`
	Finished     bool   `json:"finished"      example:"false"`
	CollectionId int    `json:"collection_id" example:"1"`
}
