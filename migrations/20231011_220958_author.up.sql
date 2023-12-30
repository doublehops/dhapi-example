CREATE TABLE author
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    name       VARCHAR(100) NOT NULL,
    created_by INT,
    updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);