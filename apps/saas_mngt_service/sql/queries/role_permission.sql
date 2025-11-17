-- name: GetRolePermissionByID :one
SELECT * FROM role_permissions WHERE id = $1 LIMIT 1;

-- name: SaveRolePermission :one
INSERT INTO role_permissions (
    id,
    user_id,
    role_id,
    permission_id,
    project_id,
    org_id,
    created_at,
    created_by_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (id) DO UPDATE
SET
    user_id = EXCLUDED.user_id,
    role_id = EXCLUDED.role_id,
    permission_id = EXCLUDED.permission_id,
    project_id = EXCLUDED.project_id,
    org_id = EXCLUDED.org_id,
    created_by_id = EXCLUDED.created_by_id
RETURNING *;
