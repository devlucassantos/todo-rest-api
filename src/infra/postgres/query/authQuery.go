package query

type authSqlManager struct{}

func Auth() *authSqlManager {
	return &authSqlManager{}
}

func (authSqlManager) Insert() string {
	return `INSERT INTO user_account (id, name, email, password)
			VALUES (DEFAULT, $1, LOWER($2), $3) RETURNING id;`
}

func (authSqlManager) UpdateToken() string {
	return `UPDATE user_account SET token = $1 WHERE id = $2;`
}

type authSelectSqlManager struct{}

func (authSqlManager) Select() *authSelectSqlManager {
	return &authSelectSqlManager{}
}

func (authSelectSqlManager) ByEmail() string {
	return `SELECT id	AS account_id,
			name 		AS account_name,
			email		AS account_email,
			password 	AS account_password
			FROM user_account WHERE email = $1;`
}
