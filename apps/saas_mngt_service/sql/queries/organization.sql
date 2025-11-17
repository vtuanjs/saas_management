-- name: GetOrganizationByID :one
SELECT * FROM organizations WHERE id = $1 LIMIT 1;

-- name: SaveOrganization :one
INSERT INTO organizations (
    id,
    code,
    name,
    logo,
    is_locked,
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
    code = EXCLUDED.code,
    name = EXCLUDED.name,
    logo = EXCLUDED.logo,
    is_locked = EXCLUDED.is_locked,
    updated_at = EXCLUDED.updated_at,
    updated_by_id = EXCLUDED.updated_by_id,
    deleted_at = EXCLUDED.deleted_at,
    version = organizations.version + 1
RETURNING *;
