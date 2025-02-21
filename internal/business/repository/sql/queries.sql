-- name: CreateBusiness :exec
INSERT INTO businesses (
    id, name, description, category, location, rating, 
    contact_info, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
);

-- name: GetBusiness :one
SELECT * FROM businesses WHERE id = $1;

-- name: UpdateBusiness :exec
UPDATE businesses 
SET name = $2, 
    description = $3, 
    category = $4,
    location = $5,
    rating = $6,
    contact_info = $7,
    updated_at = $8
WHERE id = $1;

-- name: DeleteBusiness :exec
DELETE FROM businesses WHERE id = $1;

-- name: ListBusinesses :many
SELECT * FROM businesses 
WHERE ($1::text = '' OR category = $1)
  AND ($2::text = '' OR location = $2)
LIMIT $3 OFFSET $4;