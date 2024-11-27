package response

type GenericErrorResponse struct {
	Message       string                  `json:"message"`
	InvalidFields []InvalidFieldsResponse `json:"invalid_fields,omitempty"`
	Conflicts     []string                `json:"conflicts,omitempty"`
}

type InvalidFieldsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
