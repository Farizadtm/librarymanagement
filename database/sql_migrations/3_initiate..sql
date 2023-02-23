-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    author VARCHAR(50),
    release_year INT,
    category VARCHAR(20),
    isavailable BOOLEAN NOT NULL
)

-- +migrate StatementEnd