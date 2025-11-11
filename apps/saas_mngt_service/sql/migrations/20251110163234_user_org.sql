-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_orgs" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"user_id" TEXT NOT NULL,
	"org_id" TEXT NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"version" INTEGER NOT NULL DEFAULT 1,
	"deleted_at" TIMESTAMPTZ,
	CONSTRAINT "user_orgs_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
	CONSTRAINT "user_orgs_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id")
);

CREATE UNIQUE INDEX "user_orgs_user_org_idx" ON "user_orgs" ("user_id", "org_id");
CREATE INDEX "user_orgs_org_id_idx" ON "user_orgs" ("org_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "user_orgs_org_id_idx";
DROP INDEX IF EXISTS "user_orgs_user_org_idx";
DROP TABLE IF EXISTS "user_orgs";
-- +goose StatementEnd
