-- name: GetVideo :one
SELECT * FROM video
WHERE id = $1 LIMIT 1;

-- name: CreateVideo :one
INSERT INTO video(
 src,thumbnails_count,dur,thumbnail
) VALUES ( $1,$2,$3,NULL)
RETURNING *;

-- name: UpdateVideo :one
UPDATE video
  set thumbnail= $2
WHERE id = $1
RETURNING *;
