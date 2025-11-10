-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" text PRIMARY KEY DEFAULT uuid_generate_v7(),
	"created_at" timestamptz NOT NULL DEFAULT now(),
	"updated_at" timestamptz NOT NULL DEFAULT now(),
	"created_by_id" text,
	"updated_by_id" text,
	"version" integer NOT NULL DEFAULT 1,
	"deleted_at" timestamptz,
	"email" text NOT NULL,
	"phone" text,
	"first_name" text,
	"last_name" text,
	"name" text,
	"avatar_id" text,
	"last_login" timestamptz,
	"ref_id" text,
	"is_locked" boolean NOT NULL DEFAULT false,
	"is_activated" boolean NOT NULL DEFAULT false,
	"is_admin" boolean NOT NULL DEFAULT false,
	"is_change_pass_required" boolean NOT NULL DEFAULT true
);

CREATE UNIQUE INDEX "users_email_idx" ON "users" ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "users_email_idx";
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
