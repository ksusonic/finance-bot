-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    telegram_id BIGINT       NOT NULL,
    username    VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255),
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_telegram_id UNIQUE (telegram_id)
);

CREATE TABLE transactions
(
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(255) NOT NULL,
    amount           INTEGER      NOT NULL,
    transaction_date TIMESTAMP,
    created_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id          INTEGER REFERENCES users (id),
    constraint unique_transaction UNIQUE (name, amount, transaction_date, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
