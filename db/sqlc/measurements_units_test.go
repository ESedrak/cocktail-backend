package db

import (
	"cocktail-backend/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUnit(t *testing.T) MeasurementUnit {
	randomMeasurementUnit := util.RandomNameString()
	randomIDnumber := util.RandomQty()

	args := &MeasurementUnit{
		MeasurementUnitsID: randomIDnumber,
		Unit:               randomMeasurementUnit,
	}

	measurementUnits, err := testQueries.CreateMeasurementUnit(context.Background(),
		randomMeasurementUnit,
	)

	require.NoError(t, err)
	require.NotEmpty(t, measurementUnits)

	require.Equal(t, args.Unit, measurementUnits.Unit)

	require.NotZero(t, measurementUnits.MeasurementUnitsID)

	return measurementUnits
}

func TestCreateMeasurementUnit(t *testing.T) {
	createRandomUnit(t)
}

func TestGetMeasurementUnit(t *testing.T) {
	unit1 := createRandomUnit(t)
	unit2, err := testQueries.GetMeasurementUnit(context.Background(), unit1.MeasurementUnitsID)

	require.NoError(t, err)
	require.NotEmpty(t, unit2)

	require.Equal(t, unit1.Unit, unit2.Unit)
	require.Equal(t, unit1.MeasurementUnitsID, unit2.MeasurementUnitsID)
}

func TestUpdateMeasurementUnit(t *testing.T) {
	randomMeasurementUnit := util.RandomNameString()
	unit1 := createRandomUnit(t)

	arg := UpdateMeasurementUnitsParams{
		MeasurementUnitsID: unit1.MeasurementUnitsID,
		Unit:               randomMeasurementUnit,
	}

	ingredient2, err := testQueries.UpdateMeasurementUnits(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, unit1.MeasurementUnitsID, ingredient2.MeasurementUnitsID)
	require.Equal(t, arg.Unit, ingredient2.Unit)
}

func TestDeleteMeasurementUnit(t *testing.T) {
	unit1 := createRandomUnit(t)
	err := testQueries.DeleteMeasurementUnits(context.Background(), unit1.MeasurementUnitsID)
	require.NoError(t, err)

	unit2, err := testQueries.GetMeasurementUnit(context.Background(), unit1.MeasurementUnitsID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, unit2)
}

func TestListMeasurementsUnits(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUnit(t)
	}

	arg := ListMeasurementUnitsParams{
		Limit:  5,
		Offset: 5,
	}

	units, err := testQueries.ListMeasurementUnits(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, units, 5)

	for _, unit := range units {
		require.NotEmpty(t, unit)
	}
}
