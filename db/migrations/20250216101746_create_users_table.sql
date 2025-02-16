-- migrate:up
CREATE TABLE
    IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    failed_logins INT DEFAULT 0 NOT NULL,
    locked_until BIGINT DEFAULT 0 NOT NULL
);

-- migrate:down
DROP TABLE users;
