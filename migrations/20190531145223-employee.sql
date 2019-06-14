
-- +migrate Up
CREATE TABLE employee(
    id serial PRIMARY KEY,
    name varchar,
    email varchar UNIQUE,
    password varchar,
    mobile integer,
    country varchar
);
-- +migrate Down
DROP TABLE employee;
