-- +goose Up
-- +goose StatementBegin
CREATE TABLE "organizations" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"code" TEXT NOT NULL,
	"name" TEXT NOT NULL,
	"logo" JSON,
	"is_locked" BOOLEAN NOT NULL DEFAULT FALSE,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"version" INTEGER NOT NULL DEFAULT 1,
	"deleted_at" TIMESTAMPTZ
);

CREATE UNIQUE INDEX organizations_code_idx ON organizations(code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "organizations_code_idx";
DROP TABLE IF EXISTS "organizations";
-- +goose StatementEnd
