package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Store provides all functions to execute db queries and transaction
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// CreateCocktailParams contains the input parameters to transfer information so a cocktail can be created
type CreateCocktailTxParams struct {
	RecipeID          int64 `json:"recipe_id"`
	IngredientID      int64 `json:"ingredient_id"`
	MeasurementQtyID  int64 `json:"measurement_qty_id"`
	MeasurementUnitID int64 `json:"measurement_unit_id"`
}

type CreateCocktailTxResult struct {
	Cocktail        Cocktail        `json:"cocktail"`
	Ingredient      Ingredient      `json:"ingredient"`
	MeasurementQty  MeasurementQty  `json:"measurement_qty"`
	MeasurementUnit MeasurementUnit `json:"measurement_unit"`
	Recipe          Recipe          `json:"recipe"`
}

// CreateCocktailTx performs the ability to create one complete cocktail from scratch
// It creates the ingredients/recipies/measurementQty/measurementUnits and to allow a single cocktail to be made
func (store *Store) CreateCocktailTx(ctx context.Context, arg CreateCocktailTxParams) (CreateCocktailTxResult, error) {
	var result CreateCocktailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Cocktail, err = q.CreateCocktail(ctx, CreateCocktailParams{
			RecipeID:           arg.RecipeID,
			IngredientID:       arg.IngredientID,
			MeasurementQtyID:   arg.MeasurementQtyID,
			MeasurementUnitsID: arg.MeasurementUnitID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
