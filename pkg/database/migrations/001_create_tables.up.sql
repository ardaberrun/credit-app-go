CREATE TABLE IF NOT EXISTS USERS (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50),
    email VARCHAR(50),
    hashed_password VARCHAR(1024),
    account_number SERIAL,
    balance SERIAL,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS TRANSACTIONS (
    id SERIAL PRIMARY KEY,
    user_id INT,
    amount SERIAL,
    created_at TIMESTAMP
);