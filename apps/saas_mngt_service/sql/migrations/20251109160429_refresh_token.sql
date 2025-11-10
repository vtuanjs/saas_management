-- +goose Up
-- +goose StatementBegin
CREATE TABLE "refresh_tokens" (
	"id" text PRIMARY KEY DEFAULT uuid_generate_v7(),
	"created_at" timestamptz NOT NULL DEFAULT now(),
	"updated_at" timestamptz NOT NULL DEFAULT now(),
	"token" text NOT NULL,
	"expires_at" timestamptz,
	"device_id" text,
	"device_name" text,
	"device_type" text,
	"device_os" text,
	"device_browser" text,
	"ip" text,
	"user_id" text NOT NULL,

	CONSTRAINT "refresh_tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
CREATE INDEX "refresh_tokens_user_id_idx" ON "refresh_tokens" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "refresh_tokens_user_id_idx";
DROP TABLE IF EXISTS "refresh_tokens";
-- +goose StatementEnd
