// Code generated by sqlc. DO NOT EDIT.
// source: measurement_units.sql

package db

import (
	"context"
)

const createMeasurementUnit = `-- name: CreateMeasurementUnit :one
INSERT INTO measurement_units (
  unit
) VALUES (
  $1
)
RETURNING measurement_units_id, unit
`

func (q *Queries) CreateMeasurementUnit(ctx context.Context, unit string) (MeasurementUnit, error) {
	row := q.db.QueryRowContext(ctx, createMeasurementUnit, unit)
	var i MeasurementUnit
	err := row.Scan(&i.MeasurementUnitsID, &i.Unit)
	return i, err
}

const deleteMeasurementUnits = `-- name: DeleteMeasurementUnits :exec
DELETE FROM measurement_units
WHERE measurement_units_id = $1
`

func (q *Queries) DeleteMeasurementUnits(ctx context.Context, measurementUnitsID int64) error {
	_, err := q.db.ExecContext(ctx, deleteMeasurementUnits, measurementUnitsID)
	return err
}

const getMeasurementUnit = `-- name: GetMeasurementUnit :one
SELECT measurement_units_id, unit FROM measurement_units
WHERE measurement_units_id = $1 LIMIT 1
`

func (q *Queries) GetMeasurementUnit(ctx context.Context, measurementUnitsID int64) (MeasurementUnit, error) {
	row := q.db.QueryRowContext(ctx, getMeasurementUnit, measurementUnitsID)
	var i MeasurementUnit
	err := row.Scan(&i.MeasurementUnitsID, &i.Unit)
	return i, err
}

const listMeasurementUnits = `-- name: ListMeasurementUnits :many
SELECT measurement_units_id, unit FROM measurement_units
ORDER BY measurement_units_id
LIMIT $1
OFFSET $2
`

type ListMeasurementUnitsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMeasurementUnits(ctx context.Context, arg ListMeasurementUnitsParams) ([]MeasurementUnit, error) {
	rows, err := q.db.QueryContext(ctx, listMeasurementUnits, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MeasurementUnit{}
	for rows.Next() {
		var i MeasurementUnit
		if err := rows.Scan(&i.MeasurementUnitsID, &i.Unit); err != nil {
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

const updateMeasurementUnits = `-- name: UpdateMeasurementUnits :one
UPDATE measurement_units
  set unit = $2
WHERE measurement_units_id = $1
RETURNING measurement_units_id, unit
`

type UpdateMeasurementUnitsParams struct {
	MeasurementUnitsID int64  `json:"measurement_units_id"`
	Unit               string `json:"unit"`
}

func (q *Queries) UpdateMeasurementUnits(ctx context.Context, arg UpdateMeasurementUnitsParams) (MeasurementUnit, error) {
	row := q.db.QueryRowContext(ctx, updateMeasurementUnits, arg.MeasurementUnitsID, arg.Unit)
	var i MeasurementUnit
	err := row.Scan(&i.MeasurementUnitsID, &i.Unit)
	return i, err
}
