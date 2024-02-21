CREATE TABLE IF NOT EXISTS accounts (
    id bytea NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    PRIMARY KEY(id)
);