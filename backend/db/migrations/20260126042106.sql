-- Rename a column from "is_deleted_at" to "is_deleted"
ALTER TABLE "public"."comments" RENAME COLUMN "is_deleted_at" TO "is_deleted";
-- Rename a column from "is_deleted_at" to "is_deleted"
ALTER TABLE "public"."posts" RENAME COLUMN "is_deleted_at" TO "is_deleted";
-- Rename a column from "is_deleted_at" to "is_deleted"
ALTER TABLE "public"."users" RENAME COLUMN "is_deleted_at" TO "is_deleted";
