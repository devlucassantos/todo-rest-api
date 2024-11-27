\connect todo_db;

COPY user_account (id, name, email, password)
FROM '/fixtures/user_account.csv'
CSV HEADER;
SELECT SETVAL('user_account_id_seq', (SELECT MAX(id) FROM user_account));

COPY collection (id, name, user_id)
FROM '/fixtures/collection.csv'
CSV HEADER;
SELECT SETVAL('collection_id_seq', (SELECT MAX(id) FROM collection));

COPY task (id, description, finished, user_id, collection_id)
FROM '/fixtures/task.csv'
CSV HEADER;
SELECT SETVAL('task_id_seq', (SELECT MAX(id) FROM task));