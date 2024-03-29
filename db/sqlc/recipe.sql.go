// Code generated by sqlc. DO NOT EDIT.
// source: recipe.sql

package db

import (
	"context"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipe (
  drink_name, instructions, image_url
) VALUES (
  $1, $2, $3
)
RETURNING recipe_id, drink_name, instructions, image_url, created_at
`

type CreateRecipeParams struct {
	DrinkName    string  `json:"drink_name"`
	Instructions string  `json:"instructions"`
	ImageUrl     *string `json:"image_url"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe, arg.DrinkName, arg.Instructions, arg.ImageUrl)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.DrinkName,
		&i.Instructions,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM recipe
WHERE recipe_id = $1
`

func (q *Queries) DeleteRecipe(ctx context.Context, recipeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, recipeID)
	return err
}

const getRecipe = `-- name: GetRecipe :one
SELECT recipe_id, drink_name, instructions, image_url, created_at FROM recipe
WHERE recipe_id = $1 LIMIT 1
`

func (q *Queries) GetRecipe(ctx context.Context, recipeID int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, recipeID)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.DrinkName,
		&i.Instructions,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const listRecipes = `-- name: ListRecipes :many
SELECT recipe_id, drink_name, instructions, image_url, created_at FROM recipe
ORDER BY recipe_id
LIMIT $1
OFFSET $2
`

type ListRecipesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRecipes(ctx context.Context, arg ListRecipesParams) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.RecipeID,
			&i.DrinkName,
			&i.Instructions,
			&i.ImageUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRecipe = `-- name: UpdateRecipe :one
UPDATE recipe
  set drink_name = $2,
  instructions = $3,
  image_url = $4
WHERE recipe_id = $1
RETURNING recipe_id, drink_name, instructions, image_url, created_at
`

type UpdateRecipeParams struct {
	RecipeID     int64   `json:"recipe_id"`
	DrinkName    string  `json:"drink_name"`
	Instructions string  `json:"instructions"`
	ImageUrl     *string `json:"image_url"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipe,
		arg.RecipeID,
		arg.DrinkName,
		arg.Instructions,
		arg.ImageUrl,
	)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.DrinkName,
		&i.Instructions,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}
