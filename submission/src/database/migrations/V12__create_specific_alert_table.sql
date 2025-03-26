CREATE TABLE IF NOT EXISTS specific_alert (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    breed TEXT,
    life_stage life_stage_enum,
    gender gender_enum,
    actioned BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (breed) REFERENCES dog_breed(name)
);