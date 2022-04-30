package response

type Auth struct {
	Id          int    `json:"id,omitempty"`
	AccessToken string `json:"access_token"`
}
