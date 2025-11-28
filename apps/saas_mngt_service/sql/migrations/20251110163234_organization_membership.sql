-- +goose Up
-- +goose StatementBegin
CREATE TABLE "organization_memberships" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"user_id" TEXT NOT NULL,
	"org_id" TEXT NOT NULL,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1,

	CONSTRAINT "organization_memberships_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
	CONSTRAINT "organization_memberships_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id")
);

CREATE UNIQUE INDEX "organization_memberships_org_id_user_id_idx" ON "organization_memberships" ("org_id", "user_id");
CREATE INDEX "organization_memberships_user_id_idx" ON "organization_memberships" ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "organization_memberships_org_id_user_id_idx";
DROP INDEX IF EXISTS "organization_memberships_user_id_idx";
DROP TABLE IF EXISTS "organization_memberships";
-- +goose StatementEnd
