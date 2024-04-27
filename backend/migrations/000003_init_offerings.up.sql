DO
$$
    BEGIN
        CREATE TYPE semester_type AS ENUM ('Fall', 'Spring', 'Summer 1', 'Summer 2', 'Winter');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;

-- Create the 'offering' table
CREATE TABLE IF NOT EXISTS offerings
(
    id        BIGSERIAL PRIMARY KEY,
    course_id BIGINT        NOT NULL,
    semester  semester_type NOT NULL,
    year      INTEGER       NOT NULL,
    location  TEXT          NOT NULL,
    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE,
    CONSTRAINT year_gt_1000 CHECK (year > 1000),
    CONSTRAINT year_lt_9999 CHECK (year < 9999),
    CONSTRAINT unique_offering UNIQUE (course_id, semester, year, location)
);

-- Create the 'teaching' junction table for ManyToMany relationship
CREATE TABLE IF NOT EXISTS teaches
(
    offering_id   BIGINT NOT NULL,
    instructor_id BIGINT NOT NULL,
    PRIMARY KEY (offering_id, instructor_id),
    CONSTRAINT fk_offering_id FOREIGN KEY (offering_id) REFERENCES offerings (id) ON DELETE CASCADE,
    CONSTRAINT fk_instructor_id FOREIGN KEY (instructor_id) REFERENCES instructors (id) ON DELETE CASCADE
);
