-- name: GetOrgSyncTemplateByID :one
SELECT * FROM org_sync_templates WHERE id = $1 LIMIT 1;

-- name: SaveOrgSyncTemplate :one
INSERT INTO org_sync_templates (
    id,
    template_name,
    template_type,
    template_data,
    active,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (id) DO UPDATE
SET
    template_name = EXCLUDED.template_name,
    template_type = EXCLUDED.template_type,
    template_data = EXCLUDED.template_data,
    active = EXCLUDED.active,
    updated_at = EXCLUDED.updated_at,
    version = org_sync_templates.version + 1
RETURNING *;
