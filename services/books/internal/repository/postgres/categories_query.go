package postgres

const (
	QueryCreateCategory = `
		INSERT INTO books.categories (
			id,
			name,
			description,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		)
		VALUES (
			$1,
			$2,
			$3,
			1::BIT,
			NOW(),
			$4,
			NOW(),
			$4
		)
	`

	QueryGetCategoryByID = `
		SELECT
			id,
			name,
			description
		FROM books.categories
		WHERE
			id = $1
			AND is_active = 1::BIT
	`
)
