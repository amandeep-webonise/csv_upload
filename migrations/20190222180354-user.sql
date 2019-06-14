
-- +migrate Up
CREATE TABLE users(
    id serial PRIMARY KEY,
    name varchar,
    email varchar UNIQUE,
    password varchar
);
-- +migrate Down
DROP TABLE users;
