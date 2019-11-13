package postgres

const (
	QueryGetAllExperience = `
		SELECT id, description, lokasi
		FROM experience
	`

	QueryGetExperienceByID = `
		SELECT id, description, lokasi
		FROM experience
		WHERE id = $1
	`

	QueryCreateExperience = `
		INSERT INTO experience(description, lokasi)
		VALUES ($1, $2)
		RETURNING id
	`

	QueryUpdateExperience = `
		UPDATE experience
		SET description=$2, lokasi=$3
		WHERE id=$1
	`

	QueryDeleteExperience = `
		DELETE FROM experience
		WHERE id=$1
	`
)
