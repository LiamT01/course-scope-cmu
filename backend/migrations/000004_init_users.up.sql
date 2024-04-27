-- Create the 'user' table
CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    andrew_id     VARCHAR(20)              NOT NULL UNIQUE,
    username      VARCHAR(150)             NOT NULL UNIQUE,
    password_hash BYTEA                    NOT NULL,
    admin         BOOLEAN                  NOT NULL DEFAULT false,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    avatar        TEXT, -- File paths stored as text
    activated     BOOLEAN                  NOT NULL DEFAULT false,
    CONSTRAINT andrew_id_format CHECK (andrew_id ~ '^[a-z]+[0-9]*$')
);

-- Note: Adjust data types and constraints as necessary to match Django's exact AbstractUser implementation
-- and your PostgreSQL configuration (e.g., handling of images/files and TIMESTAMP fields).
