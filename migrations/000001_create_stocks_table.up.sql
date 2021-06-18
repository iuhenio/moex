CREATE EXTENSION IF NOT EXISTS pgcrypto;

create table IF NOT EXISTS stocks(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker TEXT,
    quantity INT,
    status TEXT,
    buy_date DATE default current_date,
    buy_price DECIMAL,
    sell_date DATE NOT NULL,
    sell_price DECIMAL
);