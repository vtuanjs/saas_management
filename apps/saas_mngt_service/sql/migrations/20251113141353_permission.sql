-- +goose Up
-- +goose StatementBegin
CREATE TABLE "permissions" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"name" TEXT NOT NULL,
	"description" TEXT NOT NULL DEFAULT '',
	"system" BOOLEAN NOT NULL DEFAULT FALSE,
	"org_id" TEXT NOT NULL,
	
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1,

	CONSTRAINT "permissions_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id")
);
CREATE UNIQUE INDEX "permissions_name_idx" ON "permissions" ("name", "org_id", "deleted_at");
CREATE INDEX "permissions_org_id_idx" ON "permissions" ("org_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "permissions_name_idx";
DROP INDEX IF EXISTS "permissions_org_id_idx";
DROP TABLE IF EXISTS "permissions";
-- +goose StatementEnd
