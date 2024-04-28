-- Drop the existing unique constraint
ALTER TABLE courses DROP CONSTRAINT IF EXISTS unique_course;

-- Add a new unique constraint on the number column
ALTER TABLE courses ADD CONSTRAINT number_unique UNIQUE (number);
