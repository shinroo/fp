ALTER TABLE
    dog
ADD
    COLUMN breed TEXT;

ALTER TABLE
    dog
ADD
    FOREIGN KEY (breed) REFERENCES dog_breed(name);