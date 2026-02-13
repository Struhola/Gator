-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        gen_random_uuid(), 
        NOW(), 
        NOW(), 
        $1, -- user_id
        $2  -- feed_id
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;


-- name: GetFeedFollowsForUser :many
SELECT ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id, f.name as feed_name, u.name as user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE u.name = $1;

-- name: DeleteFeedFollow :exec
Delete from feed_follows 
where user_id = $1 and feed_id = $2;