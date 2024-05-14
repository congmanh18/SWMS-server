-- +migrate Up
CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "first_name" text,
  "middle_name" text,
  "last_name" text,
  "gender" text NOT NULL,
  "role_name" text NOT NULL,
  "date_of_birth" date,
  "email" text,
  "username" text, 
  "phone" text ,
  "password" text,
  "category" text NOT NULL,
  "token" text,
  "refresh_token" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE areas (
  "id" text PRIMARY KEY,
  "name" text,
  "address" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "trash_bins" (
  "id" text PRIMARY KEY,
  "trash_level" decimal(5,2),
  "address" text,
  "area_id" text,
  FOREIGN KEY ("area_id") REFERENCES "areas" ("id"),
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);  

CREATE TABLE "transactions" (
  "id" text PRIMARY KEY,
  "user_id" text,
  "trash_bin_id" text,
  "report_id" text,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  FOREIGN KEY ("trash_bin_id") REFERENCES "trash_bins" ("id"),
  "updated_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL
);

-- bo sung sau
CREATE TABLE permission (
    "id" text PRIMARY KEY,
    "permission_name" text
);

CREATE TABLE roles (
  "id" text PRIMARY KEY,
  "permission_id" text,
  "role_name" text UNIQUE,
  FOREIGN KEY ("permission_id") REFERENCES "permission" ("id")
);

CREATE TABLE user_roles (
  "user_id" text,
  "role_id" text,
  PRIMARY KEY ("user_id", "role_id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id")
);

-- +migrate Down
DROP TABLE user_roles;
DROP TABLE roles;
DROP TABLE permission;
DROP TABLE transactions;
DROP TABLE trash_bins;
DROP TABLE users;
DROP TABLE areas;
