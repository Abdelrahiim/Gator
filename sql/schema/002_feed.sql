-- +goose Up
CREATE TABLE "feed" (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "feed";