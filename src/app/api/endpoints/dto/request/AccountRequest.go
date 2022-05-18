package request

type Account struct {
	Name     string `json:"name"     example:"Example Name"`
	Email    string `json:"email"    example:"example@example.com"`
	Password string `json:"password" example:"ex@mplePassw0rd"`
}
