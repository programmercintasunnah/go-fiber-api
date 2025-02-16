-- migrate:up
CREATE TABLE
    IF NOT EXISTS iktikaf (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    notes TEXT
);

-- migrate:down
DROP TABLE iktikaf;