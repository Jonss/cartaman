CREATE TABLE cards(
    id bigserial PRIMARY KEY,
    suit varchar(10) NOT NULL,
    value varchar(10) NOT NULL,
    code varchar(2) UNIQUE
);