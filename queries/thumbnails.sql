-- name: GetThumb :one
SELECT * FROM thumbnails
WHERE id = $1 LIMIT 1;

-- name: InsertThumbs :one
INSERT INTO thumbnails(
 src,video,timestamp,idx
) VALUES ( $1,$2,$3,$4)
RETURNING *;

-- name: GetAllThubmsOfVideo :many
SELECT *
FROM thumbnails AS t
WHERE t.video = $1;
