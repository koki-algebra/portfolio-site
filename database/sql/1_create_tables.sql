CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- noqa: RF05

CREATE TABLE IF NOT EXISTS users (
    user_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    auth_id TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT uq_users_01 UNIQUE (auth_id),
    CONSTRAINT uq_users_02 UNIQUE (email)
);

CREATE INDEX idx_users_01 ON users (auth_id);
