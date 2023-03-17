package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCocktail(t *testing.T) Cocktail {
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

	return cocktail
}

func TestCreateCocktail(t *testing.T) {
	createRandomCocktail(t)
}

func TestGetCocktail(t *testing.T) {
	cocktail1 := createRandomCocktail(t)
	cocktail2, err := testQueries.GetCocktail(context.Background(), cocktail1.CocktailID)

	require.NoError(t, err)
	require.NotEmpty(t, cocktail2)

	require.Equal(t, cocktail1.CocktailID, cocktail2.CocktailID)
	require.Equal(t, cocktail1.RecipeID, cocktail2.RecipeID)
	require.Equal(t, cocktail1.IngredientID, cocktail2.IngredientID)
	require.Equal(t, cocktail1.MeasurementQtyID, cocktail2.MeasurementQtyID)
	require.Equal(t, cocktail1.MeasurementUnitsID, cocktail2.MeasurementUnitsID)
}

func TestUpdateCocktail(t *testing.T) {
	cocktail1 := createRandomCocktail(t)
	recipe1 := createRandomRecipe(t)
	ingredient1 := createRandomIngredient(t)
	measurementQty1 := createRandomQty(t)
	measurementUnit1 := createRandomUnit(t)

	arg := UpdateCocktailParams{
		CocktailID:   cocktail1.CocktailID,
		RecipeID:     recipe1.RecipeID,
		IngredientID: ingredient1.IngredientID,
		MeasurementQtyID: sql.NullInt64{
			Int64: measurementQty1.MeasurementQtyID,
			Valid: true,
		},
		MeasurementUnitsID: sql.NullInt64{
			Int64: measurementUnit1.MeasurementUnitsID,
			Valid: true,
		},
	}

	cocktail2, err := testQueries.UpdateCocktail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, cocktail2)

	require.Equal(t, cocktail1.CocktailID, cocktail2.CocktailID)
	require.Equal(t, ingredient1.IngredientID, cocktail2.IngredientID)
	require.Equal(t, measurementQty1.MeasurementQtyID, cocktail2.MeasurementQtyID.Int64)
	require.Equal(t, measurementUnit1.MeasurementUnitsID, cocktail2.MeasurementUnitsID.Int64)
	require.Equal(t, recipe1.RecipeID, cocktail2.RecipeID)
}

func TestDeleteCocktail(t *testing.T) {
	cocktail1 := createRandomCocktail(t)
	err := testQueries.DeleteCocktail(context.Background(), cocktail1.CocktailID)
	require.NoError(t, err)

	cocktail2, err := testQueries.GetCocktail(context.Background(), cocktail1.CocktailID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cocktail2)
}

func TestListCocktail(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCocktail(t)
	}

	arg := ListCocktailsParams{
		Limit:  5,
		Offset: 5,
	}

	cocktail, err := testQueries.ListCocktails(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, cocktail, 5)

	for _, ingredient := range cocktail {
		require.NotEmpty(t, ingredient)
	}
}
