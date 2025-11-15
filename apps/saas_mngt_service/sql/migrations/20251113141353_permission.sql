-- +goose Up
-- +goose StatementBegin
CREATE TABLE "permissions" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"code" TEXT NOT NULL,
	"name" TEXT NOT NULL,
	"description" TEXT,
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
CREATE UNIQUE INDEX "permissions_code_idx" ON "permissions" ("code", "org_id", "deleted_at");
CREATE INDEX "permissions_org_id_idx" ON "permissions" ("org_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "permissions_code_idx";
DROP INDEX IF EXISTS "permissions_org_id_idx";
DROP TABLE IF EXISTS "permissions";
-- +goose StatementEnd
