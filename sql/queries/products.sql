-- name: GetProductsByLocationId :many
SELECT * FROM stocks AS s 
JOIN products AS p ON p.product_id = s.product_id 
WHERE location_id = $1
ORDER BY p.product_id
LIMIT $2
OFFSET $3
;

-- name: GetProductsBySupplierId :many
SELECT * FROM stocks AS s 
JOIN products AS p ON p.product_id = s.product_id 
WHERE supplier_id = $1
ORDER BY p.product_id
LIMIT $2
OFFSET $3
;

-- name: GetProductById :one
SELECT * FROM products
WHERE product_id = $1;

-- name: CreateProduct :one
INSERT INTO products (product_name, description, unit_price, supplier_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET unit_price = $1, description = $2, supplier_id = $3
WHERE product_id = $4
RETURNING *;


-- SELECT * FROM stocks AS s 
-- JOIN products AS p ON p.product_id = s.product_id 
-- WHERE s.location_id = 1 
-- ORDER BY p.product_id
-- LIMIT 5
-- OFFSET 0;