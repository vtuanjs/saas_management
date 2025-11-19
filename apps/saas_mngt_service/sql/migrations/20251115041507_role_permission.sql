-- +goose Up
-- +goose StatementBegin
CREATE TABLE "role_permissions" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"user_id" TEXT NOT NULL,
	"role_id" TEXT NOT NULL,
	"permission_id" TEXT NOT NULL,
	"org_id" TEXT NOT NULL,
	"resources" TEXT[],

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,

	CONSTRAINT "role_permissions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
	CONSTRAINT "role_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles" ("id"),
	CONSTRAINT "role_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id"),
	CONSTRAINT "role_permissions_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id")
);
CREATE UNIQUE INDEX "role_permissions_unique_idx" ON "role_permissions" ("user_id", "permission_id", "role_id", "org_id");
CREATE INDEX "role_permissions_role_id_idx" ON "role_permissions" ("role_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "role_permissions_unique_idx";
DROP INDEX IF EXISTS "role_permissions_role_id_idx";
DROP TABLE IF EXISTS "role_permissions";
-- +goose StatementEnd
