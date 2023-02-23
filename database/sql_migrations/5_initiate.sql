-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE returnhistory (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(id),
    admin_id INT REFERENCES admin(id),
    book_id INT REFERENCES book(id),
    created_at TIMESTAMP
)

-- +migrate StatementEnd