-- +goose Up
-- +goose StatementBegin
CREATE TABLE "refresh_tokens" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"token" TEXT NOT NULL,
	"expires_at" TIMESTAMPTZ,
	"device_id" TEXT,
	"device_name" TEXT,
	"device_type" TEXT,
	"device_os" TEXT,
	"device_browser" TEXT,
	"ip" TEXT,
	"user_id" TEXT NOT NULL,

	CONSTRAINT "refresh_tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
CREATE INDEX "refresh_tokens_user_id_idx" ON "refresh_tokens" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "refresh_tokens_user_id_idx";
DROP TABLE IF EXISTS "refresh_tokens";
-- +goose StatementEnd
