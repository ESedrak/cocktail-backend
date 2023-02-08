-- name: CreateCocktail :one
INSERT INTO cocktail (
  recipe_id, ingredient_id, measurement_qty_id, measurement_units_id
) VALUES (
  $1, $2, $3, $4
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
  set recipe_id = $2,
    ingredient_id = $3,
    measurement_qty_id = $4,
    measurement_units_id = $5,
WHERE cocktail_id = $1
RETURNING *;

-- name: DeleteCocktail :exec
DELETE FROM cocktail
WHERE cocktail_id = $1;