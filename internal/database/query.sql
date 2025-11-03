-- name: GetAllBooks :many
select *
from books;



-- name: AddBook :exec
INSERT INTO books (title,author,publication_date,finished_date,rating)
VALUES (?,?,?,?,?);

-- name: GetBooksByAuthor :many
SELECT *
FROM books
WHERE author=?;

-- name: GetBooksByAuthorSortedByPublicationDate :many
SELECT *
FROM books
WHERE author=?
ORDER BY publication_date;



-- name: GetBooksByAuthorSortedByFinishedDate :many
SELECT *
FROM books
WHERE author=?
ORDER BY finished_date ASC;

-- name: GetBooksByAuthorSortedByFinishedDatev2 :many
SELECT *
FROM books
WHERE author=?
ORDER BY finished_date DESC;