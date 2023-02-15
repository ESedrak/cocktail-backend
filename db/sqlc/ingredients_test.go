package db

import (
	"cocktail-backend/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomIngredient(t *testing.T) Ingredient {
	randomIngredientName := util.RandomNameString()
	randomIDnumber := util.RandomQty()

	args := &Ingredient{
		IngredientID:   randomIDnumber,
		IngredientName: randomIngredientName,
	}

	ingredient, err := testQueries.CreateIngredient(context.Background(), randomIngredientName)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient)

	require.Equal(t, args.IngredientName, ingredient.IngredientName)

	require.NotZero(t, ingredient.IngredientID)

	return ingredient
}

func TestCreateIngredient(t *testing.T) {
	createRandomIngredient(t)
}

func TestGetIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)
	ingredient2, err := testQueries.GetIngredient(context.Background(), ingredient1.IngredientID)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, ingredient1.IngredientName, ingredient2.IngredientName)
	require.Equal(t, ingredient1.IngredientID, ingredient2.IngredientID)
}

func TestUpdateIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)

	arg := UpdateIngredientParams{
		IngredientID:   ingredient1.IngredientID,
		IngredientName: util.RandomNameString(),
	}

	ingredient2, err := testQueries.UpdateIngredient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, ingredient1.IngredientID, ingredient2.IngredientID)
	require.Equal(t, arg.IngredientName, ingredient2.IngredientName)
}

func TestDeleteIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)
	err := testQueries.DeleteIngredient(context.Background(), ingredient1.IngredientID)
	require.NoError(t, err)

	ingredient2, err := testQueries.GetIngredient(context.Background(), ingredient1.IngredientID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ingredient2)
}

func TestListIngredient(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIngredient(t)
	}

	arg := ListIngredientParams{
		Limit:  5,
		Offset: 5,
	}

	ingredients, err := testQueries.ListIngredient(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ingredients, 5)

	for _, ingredient := range ingredients {
		require.NotEmpty(t, ingredient)
	}
}
