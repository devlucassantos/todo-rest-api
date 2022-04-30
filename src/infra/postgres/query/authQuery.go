package query

type authSqlManager struct{}

func Auth() *authSqlManager {
	return &authSqlManager{}
}

func (authSqlManager) Insert() string {
	return `INSERT INTO user_account (id, name, email, password, hash, token)
			VALUES (DEFAULT, UPPER($1), LOWER($2), $3, $4, $5) RETURNING id;`
}

func (authSqlManager) UpdateToken() string {
	return `UPDATE user_account SET token = $1 WHERE id = $2`
}

func (authSqlManager) Select() *authSqlManager {
	return &authSqlManager{}
}

func (authSqlManager) ByEmail() string {
	return `SELECT id	AS account_id,
			name 		AS account_name,
			email		AS account_email,
			password 	AS account_password,
			hash     	AS account_hash,
			token 	 	AS account_token
			FROM user_account WHERE email = $1`
}
