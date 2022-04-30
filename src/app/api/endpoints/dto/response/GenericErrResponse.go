package response

type GenericErrResponse struct {
	Message       string                  `json:"error_message"`
	InvalidFields []InvalidFieldsResponse `json:"invalid_fields,omitempty"`
	Conflicts     []string                `json:"conflicts,omitempty"`
}

type InvalidFieldsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
