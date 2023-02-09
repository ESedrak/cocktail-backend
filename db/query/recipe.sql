-- name: CreateRecipe :one
INSERT INTO recipe (
  drink_name, instructions, image_url
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetRecipe :one
SELECT * FROM recipe
WHERE recipe_id = $1 LIMIT 1;

-- name: ListRecipes :many
SELECT * FROM recipe
ORDER BY recipe_id
LIMIT $1
OFFSET $2;

-- name: UpdateRecipe :one
UPDATE recipe
  set drink_name = $2,
  instructions = $3
WHERE recipe_id = $1
RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM recipe
WHERE recipe_id = $1;