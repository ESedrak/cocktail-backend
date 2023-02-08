-- name: CreateMeasurementQty :one
INSERT INTO measurement_qty (
  qty_amount
) VALUES (
  $1
)
RETURNING *;

-- name: GetMeasurementQty :one
SELECT * FROM measurement_qty
WHERE measurement_qty_id = $1 LIMIT 1;

-- name: ListMeasurementQty :many
SELECT * FROM measurement_qty
ORDER BY measurement_qty_id
LIMIT $1
OFFSET $2;

-- name: UpdateMeasurementQty :one
UPDATE measurement_qty
  set qty_amount = $2
WHERE measurement_qty_id = $1
RETURNING *;

-- name: DeleteMeasurementQty :exec
DELETE FROM measurement_qty
WHERE measurement_qty_id = $1;