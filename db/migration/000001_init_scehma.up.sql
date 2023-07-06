CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    deleted_at TIMESTAMPTZ,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255)
);
