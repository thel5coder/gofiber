-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "email"      varchar(50)          NOT NULL,
    "password"   varchar(128)         NOT NULL,
    "created_at" timestamp            NOT NULL,
    "updated_at" timestamp            NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "users";