-- +goose Up
-- +goose StatementBegin
CREATE TYPE "template_type" AS ENUM ('PERMISSION', 'ROLE');

CREATE TABLE "org_sync_templates" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"template_name" TEXT NOT NULL,
	"template_type" template_type NOT NULL,
	"template_data" JSON NOT NULL DEFAULT '{}'::JSON,
	"active" BOOLEAN NOT NULL DEFAULT TRUE,

	"created_at" TIMESTAMPTZ DEFAULT NOW(),
	"updated_at" TIMESTAMPTZ DEFAULT NOW(),
	"version" INTEGER NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "org_sync_templates";
DROP TYPE IF EXISTS "template_type";
-- +goose StatementEnd
