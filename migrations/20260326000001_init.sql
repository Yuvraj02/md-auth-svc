-- Auth Service initial schema (users).
-- Forward-only: do not edit after release; add a new migration instead.

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar(320) NOT NULL,
  "display_name" varchar(200) NOT NULL,
  "bio" text NOT NULL DEFAULT '',
  "password_hash" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" ("email");
