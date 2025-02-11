ALTER TABLE
    account
ADD
    CONSTRAINT account_unique_email_constraint UNIQUE (email);