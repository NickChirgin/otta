CONNECT ozon;

CREATE TABLE
    urls(
        id SERIAL NOT NULL PRIMARY KEY,
        url VARCHAR(100),
        shorturl VARCHAR(25)
    );