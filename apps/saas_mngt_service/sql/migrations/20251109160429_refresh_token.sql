-- +goose Up
-- +goose StatementBegin
CREATE TABLE "refresh_tokens" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"token_id" TEXT NOT NULL,
	"expires_at" TIMESTAMPTZ,
	"device" JSON NOT NULL DEFAULT '{}'::JSON,
	"ip" TEXT NOT NULL DEFAULT '',
	"user_id" TEXT NOT NULL,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"version" INTEGER NOT NULL DEFAULT 1,

	CONSTRAINT "refresh_tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
CREATE INDEX "refresh_tokens_user_id_idx" ON "refresh_tokens" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "refresh_tokens_user_id_idx";
DROP TABLE IF EXISTS "refresh_tokens";
-- +goose StatementEnd
