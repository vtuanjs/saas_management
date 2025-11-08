-- Create "users" table
CREATE TABLE "users" (
  "id" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_by_id" character varying NULL,
  "updated_by_id" character varying NULL,
  "version" integer NOT NULL DEFAULT 1,
  "deleted_at" timestamptz NULL,
  "email" character varying NULL,
  "salutation" character varying NOT NULL,
  "phone" character varying NULL,
  "first_name" character varying NULL,
  "last_name" character varying NULL,
  "name" character varying NULL,
  "avatar_id" character varying NULL,
  "last_login" timestamptz NULL,
  "ref_id" character varying NULL,
  "is_locked" boolean NOT NULL DEFAULT false,
  "is_activated" boolean NOT NULL DEFAULT false,
  "is_admin" boolean NOT NULL DEFAULT false,
  "is_change_pass_required" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "auth_refresh_tokens" table
CREATE TABLE "auth_refresh_tokens" (
  "id" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "refresh_token_id" character varying NOT NULL,
  "expires_at" timestamptz NULL,
  "device_id" character varying NOT NULL,
  "device_name" character varying NOT NULL,
  "device_type" character varying NOT NULL,
  "device_os" character varying NOT NULL,
  "device_browser" character varying NOT NULL,
  "ip" character varying NOT NULL,
  "user_id" character varying NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "auth_refresh_tokens_users_auth_refresh_tokens" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "auth_users" table
CREATE TABLE "auth_users" (
  "id" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "username" character varying NOT NULL,
  "password" character varying NOT NULL,
  "disabled" boolean NOT NULL DEFAULT false,
  "user_id" character varying NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "auth_users_users_auth_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "auth_users_user_id_key" to table: "auth_users"
CREATE UNIQUE INDEX "auth_users_user_id_key" ON "auth_users" ("user_id");
-- Create index "auth_users_username_key" to table: "auth_users"
CREATE UNIQUE INDEX "auth_users_username_key" ON "auth_users" ("username");
