-- name: GetPermissionByID :one
SELECT * FROM permissions WHERE id = $1 LIMIT 1;

-- name: SavePermission :one
INSERT INTO permissions (
    id,
    name,
    description,
    system,
    org_id,
    created_at,
    updated_at,
    created_by_id,
    updated_by_id,
    deleted_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (id) DO UPDATE
SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    system = EXCLUDED.system,
    org_id = EXCLUDED.org_id,
    updated_at = EXCLUDED.updated_at,
    updated_by_id = EXCLUDED.updated_by_id,
    deleted_at = EXCLUDED.deleted_at,
    version = permissions.version + 1
RETURNING *;
