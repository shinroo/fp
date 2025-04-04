CREATE TABLE IF NOT EXISTS app_session (
    token VARCHAR(255) PRIMARY KEY,
    account_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id)
);