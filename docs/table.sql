select * from player

drop table player

CREATE TABLE player (
    id serial primary key,
    name VARCHAR(15) NOT NULL,
    country VARCHAR(2) NOT NULL,
    score VARCHAR(2) NOT NULL
);


INSERT INTO player ( name, country, score) VALUES
    ( 'John Doe', 'US', '85'),
    ( 'Alice Smith', 'CA', '92'),
    ( 'Bob Johnson', 'UK', '78'),
    ( 'Emily Brown', 'AU', '95'),
    ( 'Michael Lee', 'DE', '88'),
    ( 'Emma Wilson', 'FR', '87'),
    ( 'Daniel Kim', 'KR', '96'),
    ( 'Olivia Chen', 'CN', '90'),
    ( 'James Wang', 'TW', '93'),
    ( 'Sophia Yamamoto', 'JP', '80');