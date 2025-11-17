-- name: GetRefreshTokenByID :one
SELECT * FROM refresh_tokens WHERE id = $1 LIMIT 1;

-- name: SaveRefreshToken :one
INSERT INTO refresh_tokens (
    id,
    token_id,
    expires_at,
    device,
    ip,
    user_id,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (id) DO UPDATE
SET
    token_id = EXCLUDED.token_id,
    expires_at = EXCLUDED.expires_at,
    device = EXCLUDED.device,
    ip = EXCLUDED.ip,
    user_id = EXCLUDED.user_id,
    updated_at = EXCLUDED.updated_at,
    version = refresh_tokens.version + 1
RETURNING *;
