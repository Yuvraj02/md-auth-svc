-- Make user ids plain text (matches frontend ids like "user-1") and seed the owner.

ALTER TABLE "users" ALTER COLUMN "id" TYPE text USING "id"::text;

ALTER TABLE "users" ADD COLUMN IF NOT EXISTS "username" varchar(64) NOT NULL DEFAULT '';
ALTER TABLE "users" ADD COLUMN IF NOT EXISTS "avatar" text NOT NULL DEFAULT '';
ALTER TABLE "users" ADD COLUMN IF NOT EXISTS "name" varchar(200) NOT NULL DEFAULT '';

UPDATE "users" SET "name" = "display_name" WHERE "name" = '' OR "name" IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_username" ON "users" ("username") WHERE "username" <> '';

CREATE TABLE IF NOT EXISTS "user_analytics" (
  "user_id" text PRIMARY KEY,
  "total_views" bigint NOT NULL DEFAULT 0,
  "total_posts" integer NOT NULL DEFAULT 0,
  "total_likes" bigint NOT NULL DEFAULT 0,
  "followers" integer NOT NULL DEFAULT 0,
  "views_this_month" bigint NOT NULL DEFAULT 0,
  "posts_this_month" integer NOT NULL DEFAULT 0,
  "top_posts_json" text NOT NULL DEFAULT '[]'
);

INSERT INTO "users" (
  "id", "email", "display_name", "name", "username", "avatar", "bio", "password_hash", "created_at", "updated_at"
) VALUES (
  'user-1',
  'rashi@marketingdigest.blog',
  'Rashi Koranne',
  'Rashi Koranne',
  'rashi',
  '',
  'Editor at Marketing Digest.',
  'unused',
  '2025-01-15T00:00:00Z',
  NOW()
)
ON CONFLICT ("id") DO UPDATE SET
  "email" = EXCLUDED."email",
  "display_name" = EXCLUDED."display_name",
  "name" = EXCLUDED."name",
  "username" = EXCLUDED."username",
  "bio" = EXCLUDED."bio";

INSERT INTO "user_analytics" (
  "user_id", "total_views", "total_posts", "total_likes", "followers",
  "views_this_month", "posts_this_month", "top_posts_json"
) VALUES (
  'user-1', 12840, 14, 920, 310, 2140, 2,
  '[{"id":"article-1","title":"The Architecture of Silence","views":4200},{"id":"article-2","title":"Brand voice in 2026","views":2800},{"id":"article-3","title":"A/B tests that actually ship","views":1900}]'
)
ON CONFLICT ("user_id") DO NOTHING;
