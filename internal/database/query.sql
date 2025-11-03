-- name: GetAllBooks :many
select *
from books;



-- name: AddBook :exec
INSERT INTO books (title,author,publication_date,rating)
VALUES (?,?,?,?);