package postgres

const (
	QueryInsertAttachment = `
		INSERT INTO attachment (
			type,
			content
		) VALUES (
			$1,
			$2
		)
		RETURNING attachment_id
	`

	QuerySelectAttachment = `SELECT content FROM attachment WHERE attachment_id=$1`
)
