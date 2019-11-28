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

	QueryInsertRestoran = `
		INSERT INTO restoran (
			resto_kontak,
			resto_lokasi,
			resto_name,
			attachment_id
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
	`
)
