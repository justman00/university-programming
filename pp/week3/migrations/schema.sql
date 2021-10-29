CREATE TABLE IF NOT EXISTS orders (
    id    UUID        PRIMARY KEY,
    SKU   text        NOT NULL,
    price int         NOT NULL
)