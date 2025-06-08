package postgres

const (
	QueryCreateAuthor = `
		INSERT INTO books.authors (
			id,
			name,
			bio,
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
			$2,
			NOW(),
			$2
		)
	`
)
