CREATE TABLE IF NOT EXISTS dog (
    identifier VARCHAR(200) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    gender ENUM('Male', 'Female') NOT NULL,
    life_stage ENUM('Puppy', 'Adult', 'Senior') NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    spca_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (spca_id) REFERENCES spca(id)
);