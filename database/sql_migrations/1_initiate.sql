-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    major VARCHAR(256)
)

-- +migrate StatementEnd