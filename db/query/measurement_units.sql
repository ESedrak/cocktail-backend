-- name: CreateMeasurementUnit :one
INSERT INTO measurement_units (
  unit
) VALUES (
  $1
)RETURNING *;

-- name: GetMeasurementUnit :one
SELECT * FROM measurement_units
WHERE measurement_units_id = $1 LIMIT 1;

-- name: ListMeasurementUnits :many
SELECT * FROM measurement_units
ORDER BY measurement_units_id
LIMIT $1
OFFSET $2;

-- name: UpdateMeasurementUnits :one
UPDATE measurement_units
  set unit = $2,
WHERE measurement_units_id = $1
RETURNING *;

-- name: DeleteCocktail :exec
DELETE FROM measurement_units
WHERE measurement_units_id = $1;