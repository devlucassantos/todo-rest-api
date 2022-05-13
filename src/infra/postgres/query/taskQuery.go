package query

type taskSqlManager struct{}

func Task() *taskSqlManager {
	return &taskSqlManager{}
}

func (taskSqlManager) Insert() string {
	return "INSERT INTO task (id, description, finished, collection_id, user_id) VALUES (DEFAULT, $1, $2, $3, $4) RETURNING id;"
}

func (taskSqlManager) Update() string {
	return "UPDATE task SET description = $1, finished = $2, collection_id = $3 WHERE id = $4 AND user_id = $5;"
}

func (taskSqlManager) Delete() string {
	return "DELETE FROM task WHERE id = $1 AND user_id = $2;"
}

type taskSelectSqlManager struct{}

func (taskSqlManager) Select() *taskSelectSqlManager {
	return &taskSelectSqlManager{}
}

func (taskSelectSqlManager) All() string {
	return `SELECT t.id				AS task_id,
				   t.description	AS task_description,
				   t.finished		AS task_finished,
				   c.id				AS collection_id,
				   c.name			AS collection_name
			FROM task t
			INNER JOIN collection c ON t.collection_id= c.id
			WHERE t.user_id = $1;`
}

func (taskSelectSqlManager) ById() string {
	return `SELECT t.id				AS task_id,
				   t.description	AS task_description,
				   t.finished		AS task_finished,
				   c.id				AS collection_id,
				   c.name			AS collection_name
			FROM task t
			INNER JOIN collection c ON t.collection_id= c.id
			WHERE t.id = $1 AND t.user_id = $2;`
}

func (taskSelectSqlManager) ByCollection() string {
	return `SELECT t.id				AS task_id,
				   t.description	AS task_description,
				   t.finished		AS task_finished,
				   c.id				AS collection_id,
				   c.name			AS collection_name
			FROM task t
			INNER JOIN collection c ON t.collection_id= c.id
			WHERE c.id = $1 AND t.user_id = $2;`
}
