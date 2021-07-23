

-- name: SaveTransaction :one
INSERT INTO transactions (account_id, operation_type, amount)
VALUES (@account_id::string, @operation_type::int, @amount::int)
RETURNING id;