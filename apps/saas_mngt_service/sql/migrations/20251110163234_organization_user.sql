-- +goose Up
-- +goose StatementBegin
CREATE TABLE "organization_users" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"user_id" TEXT NOT NULL,
	"org_id" TEXT NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"version" INTEGER NOT NULL DEFAULT 1,
	"deleted_at" TIMESTAMPTZ,
	CONSTRAINT "organization_users_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
	CONSTRAINT "organization_users_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX "organization_users_user_org_idx" ON "organization_users" ("user_id", "org_id");
CREATE INDEX "organization_users_user_id_idx" ON "organization_users" ("user_id");
CREATE INDEX "organization_users_org_id_idx" ON "organization_users" ("org_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "organization_users_org_id_idx";
DROP INDEX IF EXISTS "organization_users_user_id_idx";
DROP INDEX IF EXISTS "organization_users_user_org_idx";
DROP TABLE IF EXISTS "organization_users";
-- +goose StatementEnd
