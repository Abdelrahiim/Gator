
-- name: CreateFeedFollow :one

WITH inserted AS (
    INSERT INTO "feed_follows" (id, user_id, feed_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted.*,
    "feed".name AS feed_name,
    "user".name AS user_name
FROM inserted
JOIN "user" ON inserted.user_id = "user".id
JOIN "feed" ON inserted.feed_id = "feed".id;


-- name: GetFeedFollowsForUser :many
SELECT
    "feed_follows".id,
    "feed_follows".user_id,
    "feed_follows".feed_id,
    "feed_follows".created_at,
    "feed_follows".updated_at,
    "feed".name AS feed_name,
    "feed".url AS feed_url,
    "user".name AS user_name
FROM "feed_follows"
JOIN "user" ON "feed_follows".user_id = "user".id
JOIN "feed" ON "feed_follows".feed_id = "feed".id
WHERE "user".name = $1
ORDER BY "feed_follows".created_at DESC;

-- name: DeleteFeedFollow :exec
DELETE FROM "feed_follows"
WHERE user_id = $1 AND feed_id = $2;
