package db

import (
	"cocktail-backend/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQty(t *testing.T) MeasurementQty {
	randomMeasurementQty := util.RandomQty()
	randomIDnumber := util.RandomQty()

	args := &MeasurementQty{
		MeasurementQtyID: randomIDnumber,
		QtyAmount:        randomMeasurementQty,
	}

	measurementQty, err := testQueries.CreateMeasurementQty(context.Background(), randomMeasurementQty)

	require.NoError(t, err)
	require.NotEmpty(t, measurementQty)

	require.Equal(t, args.QtyAmount, measurementQty.QtyAmount)

	require.NotZero(t, measurementQty.MeasurementQtyID)

	return measurementQty
}

func TestCreateMeasurementQty(t *testing.T) {
	createRandomQty(t)
}

func TestGetMeasurementQty(t *testing.T) {
	qty1 := createRandomQty(t)
	qty2, err := testQueries.GetMeasurementQty(context.Background(), qty1.MeasurementQtyID)

	require.NoError(t, err)
	require.NotEmpty(t, qty2)

	require.Equal(t, qty1.QtyAmount, qty2.QtyAmount)
	require.Equal(t, qty1.MeasurementQtyID, qty2.MeasurementQtyID)
}

func TestUpdateMeasurementQty(t *testing.T) {
	randomMeasurementQty := util.RandomQty()
	qty1 := createRandomQty(t)

	arg := UpdateMeasurementQtyParams{
		MeasurementQtyID: qty1.MeasurementQtyID,
		QtyAmount:        randomMeasurementQty,
	}

	ingredient2, err := testQueries.UpdateMeasurementQty(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, qty1.MeasurementQtyID, ingredient2.MeasurementQtyID)
	require.Equal(t, arg.QtyAmount, ingredient2.QtyAmount)
}

func TestDeleteMeasurementQty(t *testing.T) {
	qty1 := createRandomQty(t)
	err := testQueries.DeleteMeasurementQty(context.Background(), qty1.MeasurementQtyID)
	require.NoError(t, err)

	qty2, err := testQueries.GetMeasurementQty(context.Background(), qty1.MeasurementQtyID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, qty2)
}

func TestListMeasurementsQty(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomQty(t)
	}

	arg := ListMeasurementQtyParams{
		Limit:  5,
		Offset: 5,
	}

	qtys, err := testQueries.ListMeasurementQty(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, qtys, 5)

	for _, qty := range qtys {
		require.NotEmpty(t, qty)
	}
}
