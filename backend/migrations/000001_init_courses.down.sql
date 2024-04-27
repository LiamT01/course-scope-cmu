-- Drop indices if they exist
DROP INDEX IF EXISTS course_number_idx CASCADE;
DROP INDEX IF EXISTS course_name_idx CASCADE;

-- Drop the 'courses' table
DROP TABLE IF EXISTS courses CASCADE;
