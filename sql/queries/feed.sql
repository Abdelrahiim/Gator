-- name: CreateFeed :one
INSERT INTO "feed" (id, name, url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: GetFeeds :many
SELECT "feed".*, "user".name as user_name 
FROM "feed"
JOIN "user" ON "feed".user_id = "user".id 
ORDER BY "feed".created_at DESC;

-- name: GetFeed :one
SELECT "feed".*, "user".name as user_name 
FROM "feed"
JOIN "user" ON "feed".user_id = "user".id 
WHERE "feed".url = $1;

-- name: MarkAsFetched :one
UPDATE "feed"
SET last_fetched_at = Now(), 
updated_at = Now()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM "feed"
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
