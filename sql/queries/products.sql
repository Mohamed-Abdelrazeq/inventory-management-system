-- name: GetProducts :many
SELECT * FROM "stocks" AS s 
JOIN "products" p ON "p.product_id" = "s.product_id" 
WHERE ? = ? 
LIMIT ?
OFFSET ?;

-- name: GetProdcutById :one
SELECT * FROM "products"
WHERE "product_id" = 3;

-- name: CreateProduct :many
INSERT INTO "products" ("product_name", "description", "unit_price", "supplier_id")
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateProduct :one
UPDATE "products"
SET "unit_price" = ?, "description" = ?, "supplier_id" = ?
WHERE "product_id" = ?
RETURNING *;
