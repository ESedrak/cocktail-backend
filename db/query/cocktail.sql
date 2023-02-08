-- name: CreateCocktail :one
INSERT INTO cocktail (
  drink_name, instructions
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCocktail :one
SELECT * FROM cocktail
WHERE cocktail_id = $1 LIMIT 1;

-- name: ListCocktails :many
SELECT * FROM cocktail
ORDER BY cocktail_id
LIMIT $1
OFFSET $2;

-- name: UpdateCocktail :one
UPDATE cocktail
  set drink_name = $2,
  instructions = $3
WHERE cocktail_id = $1
RETURNING *;

-- name: DeleteCocktail :exec
DELETE FROM cocktail
WHERE cocktail_id = $1;