-- -- Select All Product's Stock with Pagination
-- SELECT * FROM stocks s
-- JOIN products p ON p.product_id = s.product_id
-- WHERE p.product_id = 1
-- ORDER BY s.quantity 
-- LIMIT 5
-- OFFSET 0;

-- -- Select All Stocks in a Location with Pagination
-- SELECT * FROM stocks s
-- JOIN products p ON p.product_id = s.product_id
-- JOIN locations l ON l.location_id = s.location_id
-- WHERE s.location_id = 1
-- ORDER BY s.quantity 
-- LIMIT 5
-- OFFSET 0;


