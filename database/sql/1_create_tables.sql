CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- noqa: RF05

CREATE TABLE IF NOT EXISTS users (
    user_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT uq_users_01 UNIQUE (email)
);

CREATE INDEX idx_users_01 ON users (email);
