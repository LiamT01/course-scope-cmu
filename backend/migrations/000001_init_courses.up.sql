-- Create the 'courses' table
CREATE TABLE IF NOT EXISTS courses
(
    id          BIGSERIAL PRIMARY KEY,
    number      TEXT    NOT NULL,
    name        TEXT    NOT NULL,
    department  TEXT    NOT NULL,
    units       INTEGER NOT NULL,
    description TEXT    NOT NULL,
    CONSTRAINT units_gt_0 CHECK (units > 0),
    CONSTRAINT unique_course UNIQUE (number, name, department, units)
);

-- Create indices for 'number' and 'name'
CREATE INDEX IF NOT EXISTS course_number_idx ON courses (number);
CREATE INDEX IF NOT EXISTS course_name_idx ON courses (name);
