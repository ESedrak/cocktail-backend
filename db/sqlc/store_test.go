package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCocktailTx(t *testing.T) {
	store := NewStore(testDB)

	recipe1 := createRandomRecipe(t)
	ingredient1 := createRandomIngredient(t)
	measurementUnit1 := createRandomUnit(t)
	measurementQty1 := createRandomQty(t)

	// run with concurrent transactions
	fmt.Println(store)
	fmt.Println(recipe1)
	fmt.Println(ingredient1)
	fmt.Println(measurementUnit1)
	fmt.Println(measurementQty1)

	errs := make(chan error)
	results := make(chan CreateCocktailTxResult)

	go func() {
		result, err := store.CreateCocktailTx(context.Background(), CreateCocktailTxParams{
			RecipeID:          recipe1.RecipeID,
			IngredientID:      ingredient1.IngredientID,
			MeasurementUnitID: sql.NullInt64{Int64: measurementUnit1.MeasurementUnitsID, Valid: true},
			MeasurementQtyID:  sql.NullInt64{Int64: measurementUnit1.MeasurementUnitsID, Valid: true},
		})

		errs <- err
		results <- result
	}()

	// check results
	err := <-errs
	require.NoError(t, err)

	result := <-results
	require.NotEmpty(t, result)

	cocktail := result.Cocktail
	require.NotEmpty(t, cocktail)
	require.Equal(t, recipe1.RecipeID, cocktail.RecipeID)
	require.Equal(t, ingredient1.IngredientID, cocktail.IngredientID)
	require.Equal(t, measurementQty1.MeasurementQtyID, cocktail.MeasurementQtyID.Int64)
	require.Equal(t, measurementUnit1.MeasurementUnitsID, cocktail.MeasurementUnitsID.Int64)
	require.NotZero(t, cocktail.CocktailID)

	_, err = store.GetCocktail(context.Background(), cocktail.CocktailID)
	require.NoError(t, err)
}
