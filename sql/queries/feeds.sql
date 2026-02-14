-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, last_fetched_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;


-- name: GetFeeds :many
SELECT f.id, f.created_at, f.updated_at, f.last_fetched_at, f.name, f.url, f.user_id, u.name as user_name
FROM feeds f
JOIN users u
    ON f.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT f.id, f.created_at, f.updated_at, f.last_fetched_at, f.name, f.url, f.user_id, u.name as user_name
FROM feeds f
JOIN users u
    ON f.url = $1;


-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2
    ,last_fetched_at = $2
Where id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, name, url 
FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST 
FOR UPDATE SKIP LOCKED 
LIMIT 1;