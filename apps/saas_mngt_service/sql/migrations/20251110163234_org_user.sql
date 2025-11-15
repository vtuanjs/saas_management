-- +goose Up
-- +goose StatementBegin
CREATE TABLE "org_users" (
	"id" TEXT PRIMARY KEY DEFAULT uuid_generate_v7(),
	"user_id" TEXT NOT NULL,
	"org_id" TEXT NOT NULL,

	"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"created_by_id" TEXT,
	"updated_by_id" TEXT,
	"deleted_at" TIMESTAMPTZ,
	"version" INTEGER NOT NULL DEFAULT 1,

	CONSTRAINT "org_users_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
	CONSTRAINT "org_users_org_id_fkey" FOREIGN KEY ("org_id") REFERENCES "organizations" ("id")
);

CREATE UNIQUE INDEX "org_users_org_id_user_id_idx" ON "org_users" ("org_id", "user_id");
CREATE INDEX "org_users_user_id_idx" ON "org_users" ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "org_users_org_id_user_id_idx";
DROP INDEX IF EXISTS "org_users_user_id_idx";
DROP TABLE IF EXISTS "org_users";
-- +goose StatementEnd
