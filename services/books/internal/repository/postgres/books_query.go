package postgres

const (
	QueryCreateBook = `
		INSERT INTO books.books (
			id,
			title,
			isbn,
			stock,
			author_id,
			category_id,
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
			$4,
			$5,
			$6,
			$7,
			1::BIT,
			NOW(),
			$8,
			NOW(),
			$8
		)
	`

	QueryGetBookByID = `		
		SELECT 
			b.id AS book_id,
			b.title,
			b.isbn,
			b.stock,
			b.description,
			a.id AS author_id,
			a."name" AS author_name,
			a.bio AS author_bio,
			c.id AS category_id,
			c."name" AS category_name,
			c.description AS category_description
		FROM books.books b 
		JOIN books.categories c 
			ON b.category_id = c.id AND c.is_active = 1::BIT
		JOIN books.authors a 
			ON b.author_id = a.id AND a.is_active = 1::BIT
		WHERE
			b.is_active = 1::BIT
			AND b.id = $1
	`

	QuerySearchBooks = `		
		SELECT 
			b.id AS book_id,
			b.title,
			b.isbn,
			b.stock,
			b.description,
			a.id AS author_id,
			a."name" AS author_name,
			a.bio AS author_bio,
			c.id AS category_id,
			c."name" AS category_name,
			c.description AS category_description
		FROM books.books b 
		JOIN books.categories c 
			ON b.category_id = c.id AND c.is_active = 1::BIT
		JOIN books.authors a 
			ON b.author_id = a.id AND a.is_active = 1::BIT
		WHERE
			b.is_active = 1::BIT
	`
)
