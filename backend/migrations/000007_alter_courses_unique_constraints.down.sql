-- Remove the unique constraint added on the number column
ALTER TABLE courses DROP CONSTRAINT IF EXISTS number_unique;

-- Re-add the original unique constraint
ALTER TABLE courses ADD CONSTRAINT unique_course UNIQUE (number, name, department, units);
