-- migrate:up
CREATE TABLE
    IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    description TEXT
);

-- migrate:down
DROP TABLE books;
