-- name: CreateIngredient :one
INSERT INTO ingredients (
  ingredient_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetIngredient :one
SELECT * FROM ingredients
WHERE ingredient_id = $1 LIMIT 1;

-- name: ListIngredient :many
SELECT * FROM ingredients
ORDER BY ingredient_id
LIMIT $1
OFFSET $2;

-- name: UpdateIngredient :one
UPDATE ingredients
  set ingredient_name = $2
WHERE ingredient_id = $1
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM ingredients
WHERE ingredient_id = $1;