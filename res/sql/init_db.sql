CREATE DATABASE todo_db;

\connect todo_db;

CREATE TABLE user_account
(
     id       SERIAL       PRIMARY KEY,
     name     VARCHAR(50)  NOT NULL,
     email    VARCHAR(50)  UNIQUE,
     password VARCHAR(200) NOT NULL,
     hash     VARCHAR(600) NOT NULL,
     token    VARCHAR(300) NOT NULL
);

CREATE TABLE collection
(
    id   SERIAL      PRIMARY KEY,
    name VARCHAR(50) NOT NULL,

    user_id INT NOT NULL,

    CONSTRAINT collection_user_fk FOREIGN KEY (user_id) REFERENCES user_account (id)
);

CREATE TABLE task
(
    id          SERIAL      PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    finished    BOOLEAN     NOT NULL,

    user_id       INT NOT NULL,
    collection_id INT,

    CONSTRAINT task_user_fk       FOREIGN KEY (user_id)       REFERENCES user_account (id),
    CONSTRAINT task_collection_fk FOREIGN KEY (collection_id) REFERENCES collection   (id)
);

CREATE DATABASE todo_db_test WITH TEMPLATE todo_db OWNER postgres;