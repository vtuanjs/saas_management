-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: SaveUser :one
INSERT INTO users (
    id,
    email,
    phone,
    first_name,
    last_name,
    avatar,
    last_login,
    ref,
    is_locked,
    is_activated,
    is_change_pass_required,
    created_at,
    updated_at,
    created_by_id,
    updated_by_id,
    deleted_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
)
ON CONFLICT (id) DO UPDATE
SET
    email = EXCLUDED.email,
    phone = EXCLUDED.phone,
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    avatar = EXCLUDED.avatar,
    last_login = EXCLUDED.last_login,
    ref = EXCLUDED.ref,
    is_locked = EXCLUDED.is_locked,
    is_activated = EXCLUDED.is_activated,
    is_change_pass_required = EXCLUDED.is_change_pass_required,
    updated_at = EXCLUDED.updated_at,
    updated_by_id = EXCLUDED.updated_by_id,
    deleted_at = EXCLUDED.deleted_at,
    version = users.version + 1
RETURNING *;