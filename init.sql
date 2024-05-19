CREATE TABLE IF NOT EXISTS spenders (
  "spender_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "email" string,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS transactions (
    "transaction_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
    "date" TIMESTAMP NOT NULL,
    "amount" float,
    "category" int,
    "transaction_type" int,
    "note" string,
    "image_url" string,
    "spender_id" int,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS categories (
    "category_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
    "category" string,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "used_flg" bool DEFAULT true
);

CREATE TABLE IF NOT EXISTS transaction_types (
    "transaction_type_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
    "type" string,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "used_flg" bool DEFAULT true
);
