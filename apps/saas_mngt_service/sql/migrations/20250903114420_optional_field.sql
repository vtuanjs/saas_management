-- Modify "auth_refresh_tokens" table
ALTER TABLE "auth_refresh_tokens" ALTER COLUMN "device_id" DROP NOT NULL, ALTER COLUMN "device_name" DROP NOT NULL, ALTER COLUMN "device_type" DROP NOT NULL, ALTER COLUMN "device_os" DROP NOT NULL, ALTER COLUMN "device_browser" DROP NOT NULL, ALTER COLUMN "ip" DROP NOT NULL;
