DROP TYPE IF EXISTS gender_enum;

CREATE TYPE gender_enum AS ENUM ('Male', 'Female');

DROP TYPE IF EXISTS life_stage_enum;

CREATE TYPE life_stage_enum AS ENUM ('Puppy', 'Adult', 'Senior');

CREATE TABLE IF NOT EXISTS dog (
    identifier VARCHAR(200) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    gender gender_enum NOT NULL,
    life_stage life_stage_enum NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    spca_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (spca_id) REFERENCES spca(id)
);