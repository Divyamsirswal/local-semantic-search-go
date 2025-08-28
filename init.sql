CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE sentences (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    embedding vector(1024)
);
