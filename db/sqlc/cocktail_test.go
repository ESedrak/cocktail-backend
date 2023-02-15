package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCocktail(t *testing.T) {
	recipe := createRandomRecipe(t)
	Ingredient := createRandomIngredient(t)
	unit := createRandomUnit(t)
	qty := createRandomQty(t)

	arg := CreateCocktailParams{
		RecipeID:     recipe.RecipeID,
		IngredientID: Ingredient.IngredientID,
		MeasurementUnitsID: sql.NullInt64{
			Int64: unit.MeasurementUnitsID,
			Valid: true,
		},
		MeasurementQtyID: sql.NullInt64{
			Int64: qty.MeasurementQtyID,
			Valid: true,
		},
	}

	cocktail, err := testQueries.CreateCocktail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, cocktail)

	require.Equal(t, cocktail.RecipeID, arg.RecipeID)
	require.Equal(t, cocktail.IngredientID, arg.IngredientID)
	require.Equal(t, cocktail.MeasurementUnitsID, arg.MeasurementUnitsID)
	require.Equal(t, cocktail.MeasurementQtyID, arg.MeasurementQtyID)
}

func TestCreateCocktail(t *testing.T) {
	createRandomCocktail(t)
}
