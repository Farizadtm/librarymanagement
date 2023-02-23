-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    username VARCHAR(16),
    pass VARCHAR(256),
    birth_date VARCHAR(256)
)

-- +migrate StatementEnd