CREATE TABLE authors
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    biography  TEXT,
    birth_date DATE
);

CREATE TABLE books
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(100),
    author_id INT REFERENCES authors (id),
    year      INT,
    isbn      VARCHAR(20)
);