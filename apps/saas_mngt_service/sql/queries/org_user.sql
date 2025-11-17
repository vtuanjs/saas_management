-- name: GetOrgUserByID :one
SELECT * FROM org_users WHERE id = $1 LIMIT 1;

-- name: SaveOrgUser :one
INSERT INTO org_users (
    id,
    user_id,
    org_id,
    created_at,
    updated_at,
    created_by_id,
    updated_by_id,
    deleted_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (id) DO UPDATE
SET
    user_id = EXCLUDED.user_id,
    org_id = EXCLUDED.org_id,
    updated_at = EXCLUDED.updated_at,
    updated_by_id = EXCLUDED.updated_by_id,
    deleted_at = EXCLUDED.deleted_at,
    version = org_users.version + 1
RETURNING *;
