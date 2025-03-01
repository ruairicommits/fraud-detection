-- setup.sql
CREATE DATABASE transactions;

\c transactions

CREATE TABLE IF NOT EXISTS public.transactions (
    id SERIAL PRIMARY KEY,
    transaction_id UUID NOT NULL,
    user_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    timestamp BIGINT NOT NULL
);

