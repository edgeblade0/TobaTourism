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

	QuerySelectRestoran = `
		SELECT resto_id,resto_kontak,resto_lokasi,resto_name,attachment_id
		FROM restoran
	`

	QueryUpdateImageRestoran = `
		UPDATE restoran
		SET attachment_id = $1
		WHERE resto_id = $2
	`

	QueryUpdateRestoran = `
		UPDATE restoran
		SET 
			resto_kontak = $1,
			resto_lokasi = $2,
			resto_name = $3
		WHERE resto_id = $4
	`

	QueryDeleteRestoran = `
		DELETE FROM restoran
		WHERE resto_id  = $1
	`
	QuerySelectRestoranByID = `
		SELECT resto_id,resto_kontak,resto_lokasi,resto_name,attachment_id
		FROM restoran WHERE resto_id = $1
	`
)
