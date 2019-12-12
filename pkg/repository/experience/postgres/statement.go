package postgres

const (
	QueryGetAllExperience = `
		SELECT experience_id, experience_description, experience_location, attachment_id
		FROM experience
	`

	QueryGetExperienceByID = `
		SELECT experience_id, experience_description, experience_location, attachment_id
		FROM experience
		WHERE experience_id = $1
	`

	QueryCreateExperience = `
		INSERT INTO experience(experience_description, experience_location, attachment_id)
		VALUES ($1, $2, $3)
		RETURNING experience_id
	`

	QueryUpdateExperience = `
		UPDATE experience
		SET experience_description=$2, experience_location=$3, attachment_id=$4
		WHERE experience_id=$1
	`

	QueryDeleteExperience = `
		DELETE FROM experience
		WHERE experience_id=$1
	`
)
