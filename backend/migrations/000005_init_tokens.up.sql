DO
$$
    BEGIN
        CREATE TYPE scope_type AS ENUM ('ACT', 'AUTH', 'PWD');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;

-- Create the 'token' table
CREATE TABLE IF NOT EXISTS tokens
(
    id         BIGSERIAL PRIMARY KEY,
    hash       BYTEA                    NOT NULL UNIQUE,
    user_id    BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    expiry     TIMESTAMP WITH TIME ZONE NOT NULL,
    scope      scope_type               NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Note: Adjust TIMESTAMP handling and BYTEA for binary fields as necessary to match your database configuration.
