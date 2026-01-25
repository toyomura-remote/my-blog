-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "did" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "is_deleted_at" boolean NOT NULL DEFAULT false,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create "posts" table
CREATE TABLE "public"."posts" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "did" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "content" text NOT NULL,
  "is_deleted_at" boolean NOT NULL DEFAULT false,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_posts_did" to table: "posts"
CREATE UNIQUE INDEX "idx_posts_did" ON "public"."posts" ("did");
-- Create "comments" table
CREATE TABLE "public"."comments" (
  "id" serial NOT NULL,
  "user_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
  "content" text NOT NULL,
  "is_deleted_at" boolean NOT NULL DEFAULT false,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_post_id" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
