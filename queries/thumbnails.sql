-- name: GetThumb :one
SELECT * FROM Thumbnails
WHERE id = $1 LIMIT 1;

-- name: InsertThumbs :one
INSERT INTO Thumbnails(
 src,video,timestamp,idx
) VALUES ( $1,$2,$3,$4)
RETURNING *;

-- name: GetAllThubmsOfVideo :many
SELECT *
FROM Thumbnails AS t
WHERE t.video = $1;
