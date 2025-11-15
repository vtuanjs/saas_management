-- +goose Up
-- +goose StatementBegin
CREATE TABLE "roles" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"name" TEXT NOT NULL,
	"description" TEXT,
	"system" BOOLEAN NOT NULL DEFAULT FALSE,
	"org_id" TEXT NOT NULL,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1
);
CREATE INDEX "roles_org_id_idx" ON "roles" ("org_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "roles_org_id_idx";
DROP TABLE IF EXISTS "roles";
-- +goose StatementEnd
