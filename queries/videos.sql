-- name: GetVideo :one
SELECT * FROM Video
WHERE id = $1 LIMIT 1;

-- name: CreateVideo :one
INSERT INTO Video(
 src,thumbnails_count,dur
) VALUES ( $1,$2,$3)
RETURNING *;

-- name: UpdateVideo :one
UPDATE Video
  set thumbnail= $2
WHERE id = $1
RETURNING *;
