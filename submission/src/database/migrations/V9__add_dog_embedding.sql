CREATE EXTENSION vector;

ALTER TABLE
    dog
ADD
    COLUMN embedding vector(15);