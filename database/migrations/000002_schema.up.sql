CREATE TABLE IF NOT EXISTS books (
    id int8 GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL
);

INSERT INTO public.books
(title, author, "year")
VALUES('How to read a book', 'Mortimer J. Adler', '1972');