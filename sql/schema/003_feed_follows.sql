-- +goose Up
CREATE table feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES "feed" (id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE "feed_follows";