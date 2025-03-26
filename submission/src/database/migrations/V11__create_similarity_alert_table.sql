CREATE TABLE IF NOT EXISTS similarity_alert (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    similarity_threshold DECIMAL(5, 2) NOT NULL,
    actioned BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id)
);