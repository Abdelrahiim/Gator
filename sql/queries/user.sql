-- name: CreateUser :one
INSERT INTO "user" (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user" WHERE name = $1;


-- name: DeleteUsers :exec
DELETE FROM "user";

-- name: GetUsers :many
SELECT * FROM "user" ORDER BY created_at DESC;