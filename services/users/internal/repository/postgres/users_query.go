package postgres

const (
	QueryCreateUser = `
		INSERT INTO users.users (
			id,
			email,
			name,
			password_hash,
			role,
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
			1::BIT,
			NOW(),
			$2,
			NOW(),
			$2
		)
	`

	QueryGetUserByEmail = `
		SELECT
			id,
			email,
			name,
			password_hash,
			role
		FROM users.users
		WHERE
			is_active = 1::BIT
			AND email = $1
	`
)
