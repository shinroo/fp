CREATE TABLE IF NOT EXISTS profile (
    account_id INT NOT NULL,
    has_children BOOLEAN NOT NULL,
    has_active_lifestyle BOOLEAN NOT NULL,
    has_lots_of_time BOOLEAN NOT NULL,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    nearest_spca VARCHAR(50) NOT NULL,
    FOREIGN KEY (nearest_spca) REFERENCES spca(id),
    FOREIGN KEY (account_id) REFERENCES account(id)
);