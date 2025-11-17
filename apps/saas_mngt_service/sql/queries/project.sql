-- name: GetProjectByID :one
SELECT * FROM projects WHERE id = $1 LIMIT 1;

-- name: SaveProject :one
INSERT INTO projects (
    id,
    name,
    code,
    parent_id,
    logo,
    org_id,
    created_at,
    updated_at,
    created_by_id,
    updated_by_id,
    deleted_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
ON CONFLICT (id) DO UPDATE
SET
    name = EXCLUDED.name,
    code = EXCLUDED.code,
    parent_id = EXCLUDED.parent_id,
    logo = EXCLUDED.logo,
    org_id = EXCLUDED.org_id,
    updated_at = EXCLUDED.updated_at,
    updated_by_id = EXCLUDED.updated_by_id,
    deleted_at = EXCLUDED.deleted_at,
    version = projects.version + 1
RETURNING *;
