-- +goose Up
-- +goose StatementBegin
CREATE TABLE "projects" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"name" TEXT NOT NULL,
	"code" TEXT NOT NULL,
	"parent_id" TEXT,
	"logo" JSON NOT NULL DEFAULT '{}'::JSON,
	"org_id" TEXT NOT NULL,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1,

	CONSTRAINT "projects_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id"),
	CONSTRAINT "projects_parent_id_fkey" FOREIGN KEY ("parent_id") REFERENCES "projects" ("id")
);

CREATE UNIQUE INDEX "projects_code_idx" ON "projects" ("code", "org_id", "deleted_at");
CREATE INDEX "projects_org_id_idx" ON "projects" ("org_id");
CREATE INDEX "projects_parent_id_idx" ON "projects" ("parent_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "projects_code_idx";
DROP INDEX IF EXISTS "projects_org_id_idx";
DROP INDEX IF EXISTS "projects_parent_id_idx";
DROP TABLE IF EXISTS "projects";
-- +goose StatementEnd