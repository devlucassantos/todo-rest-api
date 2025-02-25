package query

type collectionSqlManager struct{}

func Collection() *collectionSqlManager {
	return &collectionSqlManager{}
}

func (collectionSqlManager) Insert() string {
	return "INSERT INTO  collection (id, name, user_id)  VALUES (DEFAULT, $1, $2) RETURNING id;"
}

func (collectionSqlManager) Update() string {
	return "UPDATE collection SET name = $1 WHERE id = $2 AND user_id = $3;"
}

func (collectionSqlManager) Delete() string {
	return "DELETE FROM collection WHERE id = $1 AND user_id = $2;"
}

type collectionSelectSqlManager struct{}

func (collectionSqlManager) Select() *collectionSelectSqlManager {
	return &collectionSelectSqlManager{}
}

func (collectionSelectSqlManager) All() string {
	return `SELECT id   AS collection_id,
				   name AS collection_name
			FROM collection WHERE user_id = $1;`
}
