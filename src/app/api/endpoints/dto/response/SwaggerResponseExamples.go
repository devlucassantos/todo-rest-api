package response

type SwaggerAuthResponse struct {
	Id          int    `json:"id"           example:"1"`
	Name        string `json:"name"         example:"Example Name"`
	Email       string `json:"email"        example:"example@example.com"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6Ikp..."`
}

type SwaggerIdResponse struct {
	Id int `json:"id" example:"1"`
}

type SwaggerCollectionResponse struct {
	Id   int    `json:"id"   example:"1"`
	Name string `json:"name" example:"Collection example"`
}

type SwaggerTaskResponse struct {
	Id          int                        `json:"id"          example:"1"`
	Description string                     `json:"description" example:"Description example"`
	Finished    bool                       `json:"finished"    example:"false"`
	Collection  *SwaggerCollectionResponse `json:"collection"`
}

type SwaggerGenericErrorResponse struct {
	Message string `json:"error_msg" example:"Oops! An unexpected error has occurred."`
}

type SwaggerNotFoundErrorResponse struct {
	Message string `json:"error_msg" example:"Not Found"`
}

type SwaggerUnauthorizedResponse struct {
	Message string `json:"message" example:"Oops! You are not authorized."`
}

type SwaggerForbiddenResponse struct {
	Message string `json:"message" example:"Oops! You do not have access to this information."`
}

type SwaggerValidationErrorResponse struct {
	Message       string                `json:"error_msg" example:"Some of the data entered is invalid."`
	InvalidFields []SwaggerInvalidField `json:"invalid_fields"`
}

type SwaggerInvalidField struct {
	Name        string `json:"name" example:"Field example"`
	Description string `json:"description" example:"Description example"`
}

type SwaggerConflictErrorResponse struct {
	Message   string   `json:"error_msg" example:"It is not possible to perform the operation because there are conflicting and/or duplicate data."`
	Conflicts []string `json:"conflicts" example:"Field example"`
}
