\connect ozon;

CREATE TABLE
    urls(
        id SERIAL NOT NULL PRIMARY KEY,
url VARCHAR(100) UNIQUE,
shorturl VARCHAR(25) UNIQUE
    );