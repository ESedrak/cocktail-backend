package db

import (
	"cocktail-backend/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomRecipe(t *testing.T) Recipe {
	arg := CreateRecipeParams{
		DrinkName:    util.RandomNameString(),
		Instructions: util.RandomNameString(),
		ImageUrl: sql.NullString{
			String: util.RandomNameString(),
			Valid:  true,
		},
	}

	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)

	require.Equal(t, arg.DrinkName, recipe.DrinkName)
	require.Equal(t, arg.Instructions, recipe.Instructions)
	require.Equal(t, arg.ImageUrl, recipe.ImageUrl)

	require.NotZero(t, recipe.RecipeID)
	require.NotZero(t, recipe.CreatedAt)

	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}

func TestGetRecipe(t *testing.T) {
	recipe1 := createRandomRecipe(t)
	recipe2, err := testQueries.GetRecipe(context.Background(), recipe1.RecipeID)
	require.NoError(t, err)
	require.NotEmpty(t, recipe2)

	require.Equal(t, recipe1.RecipeID, recipe2.RecipeID)
	require.Equal(t, recipe1.DrinkName, recipe2.DrinkName)
	require.Equal(t, recipe1.ImageUrl, recipe2.ImageUrl)
	require.Equal(t, recipe1.Instructions, recipe2.Instructions)
	require.WithinDuration(t, recipe1.CreatedAt, recipe2.CreatedAt, time.Second)
}

func TestUpdateRecipe(t *testing.T) {
	recipe1 := createRandomRecipe(t)

	arg := UpdateRecipeParams{
		RecipeID:     recipe1.RecipeID,
		DrinkName:    util.RandomNameString(),
		Instructions: util.RandomNameString(),
		ImageUrl: sql.NullString{
			String: util.RandomNameString(),
			Valid:  true,
		},
	}

	recipe2, err := testQueries.UpdateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe2)

	require.Equal(t, recipe1.RecipeID, recipe2.RecipeID)
	require.Equal(t, arg.DrinkName, recipe2.DrinkName)
	require.Equal(t, arg.Instructions, recipe2.Instructions)
	require.Equal(t, arg.ImageUrl, recipe2.ImageUrl)
	require.WithinDuration(t, recipe1.CreatedAt, recipe2.CreatedAt, time.Second)
}

func TestDeleteRecipe(t *testing.T) {
	recipe1 := createRandomRecipe(t)
	err := testQueries.DeleteRecipe(context.Background(), recipe1.RecipeID)
	require.NoError(t, err)

	recipe2, err := testQueries.GetRecipe(context.Background(), recipe1.RecipeID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, recipe2)
}

func TestListRecipe(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRecipe(t)
	}

	arg := ListRecipesParams{
		Limit:  5,
		Offset: 5,
	}

	recipes, err := testQueries.ListRecipes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, recipes, 5)

	for _, recipe := range recipes {
		require.NotEmpty(t, recipe)
	}
}
