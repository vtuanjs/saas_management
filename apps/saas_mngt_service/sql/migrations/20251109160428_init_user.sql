-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"email" TEXT NOT NULL,
	"phone" TEXT NOT NULL,
	"first_name" TEXT NOT NULL,
	"last_name" TEXT NOT NULL,
	"name" TEXT GENERATED ALWAYS AS ("first_name" || ' ' || "last_name") STORED,
	"avatar" JSON,
	"last_login" TIMESTAMPTZ,
	"ref" TEXT,
	"is_locked" BOOLEAN NOT NULL DEFAULT FALSE,
	"is_activated" BOOLEAN NOT NULL DEFAULT FALSE,
	"is_admin" BOOLEAN NOT NULL DEFAULT FALSE,
	"is_change_pass_required" BOOLEAN NOT NULL DEFAULT TRUE,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1
);

CREATE UNIQUE INDEX "users_email_idx" ON "users" ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "users_email_idx";
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
