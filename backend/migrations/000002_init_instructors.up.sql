-- Create the 'instructor' table
CREATE TABLE IF NOT EXISTS instructors
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Create index for 'name'
CREATE INDEX IF NOT EXISTS instructor_name_idx ON instructors (name);
