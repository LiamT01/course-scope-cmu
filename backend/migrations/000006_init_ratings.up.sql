-- Create the 'rating' table
CREATE TABLE IF NOT EXISTS ratings
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT                   NOT NULL,
    offering_id BIGINT                   NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    overall     INTEGER                  NOT NULL,
    teaching    INTEGER                  NOT NULL,
    materials   INTEGER                  NOT NULL,
    value       INTEGER                  NOT NULL,
    difficulty  INTEGER                  NOT NULL,
    workload    INTEGER                  NOT NULL,
    grading     INTEGER                  NOT NULL,
    comment     TEXT                     NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_offering_id FOREIGN KEY (offering_id) REFERENCES offerings (id) ON DELETE CASCADE,
    CONSTRAINT overall_range CHECK (overall BETWEEN 1 AND 5),
    CONSTRAINT teaching_range CHECK (teaching BETWEEN 1 AND 5),
    CONSTRAINT materials_range CHECK (materials BETWEEN 1 AND 5),
    CONSTRAINT value_range CHECK (value BETWEEN 1 AND 5),
    CONSTRAINT difficulty_range CHECK (difficulty BETWEEN 1 AND 5),
    CONSTRAINT workload_range CHECK (workload BETWEEN 1 AND 5),
    CONSTRAINT grading_range CHECK (grading BETWEEN 1 AND 5)
);

-- Create the 'like' table
CREATE TABLE IF NOT EXISTS likes
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT                   NOT NULL,
    rating_id  BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_rating_id FOREIGN KEY (rating_id) REFERENCES ratings (id) ON DELETE CASCADE,
    CONSTRAINT unique_like UNIQUE (user_id, rating_id)
);

-- Create the 'dislike' table
CREATE TABLE IF NOT EXISTS dislikes
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT                   NOT NULL,
    rating_id  BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_rating_id FOREIGN KEY (rating_id) REFERENCES ratings (id) ON DELETE CASCADE,
    CONSTRAINT unique_dislike UNIQUE (user_id, rating_id)
);
