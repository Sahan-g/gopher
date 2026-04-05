-- +goose Up

-- enable extension (safe to run multiple times)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- set default value for id
ALTER TABLE users
ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- +goose Down

ALTER TABLE users
ALTER COLUMN id DROP DEFAULT;