-- name: GetAllBooks :many
select *
from books;


-- name: GetAllBooksSortedByDate :many
select *
from books
ORDER BY date(finished_date) DESC;





-- name: AddBook :exec
INSERT OR IGNORE INTO books (title,author,finished_date,rating,uri_thumbnail,review)
VALUES (?,?,?,?,?,?);

-- name: GetBooksByAuthor :many
SELECT *
FROM books
WHERE author=?;






-- name: GetBooksSortedByFinishedDateAsc :many
SELECT *
FROM books
ORDER BY finished_date ASC;

-- name: GetBooksSortedByFinishedDateDesc :many
SELECT *
FROM books
ORDER BY finished_date DESC;


-- name: GetBookById :one
SELECT *
FROM books
WHERE id=?
LIMIT 1;

-- name: DeleteBookById :exec
DELETE FROM books
where id=?;



-- name: UpdateReviewById :one
UPDATE books
SET review=?
WHERE id=?
RETURNING *;

-- name: UpdateRatingById :one
UPDATE books
SET rating=?
WHERE id=?
RETURNING *;
