schema "public" {}

table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "did" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "email" {
    null = false
    type = varchar(255)
  }
  column "password" {
    null = false
    type = varchar(255)
  }
  column "is_deleted_at" {
    null    = false
    type    = boolean
    default = false
  }
  column "updated_at" {
    null    = true
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_users_email" {
    unique  = true
    columns = [column.email]
  }
}

table "posts" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "user_id" {
    null = false
    type = bigint
  }
  column "did" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "content" {
    null = false
    type = text
  }
  column "is_deleted_at" {
    null    = false
    type    = boolean
    default = false
  }
  column "updated_at" {
    null    = true
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_posts_did" {
    unique  = true
    columns = [column.did]
  }
  foreign_key "fk_user_id" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}

table "comments" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "user_id" {
    null = false
    type = bigint
  }
  column "post_id" {
    null = false
    type = bigint
  }
  column "content" {
    null = false
    type = text
  }
  column "is_deleted_at" {
    null    = false
    type    = boolean
    default = false
  }
  column "updated_at" {
    null    = true
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_user_id" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
  foreign_key "fk_post_id" {
    columns = [column.post_id]
    ref_columns = [table.posts.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}