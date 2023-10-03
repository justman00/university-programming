CREATE TABLE entities (
    id UUID PRIMARY KEY,
    external_id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
