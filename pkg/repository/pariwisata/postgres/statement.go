package postgres

const (
	QueryGetAllPariwisata = `
		SELECT pariwisata_id, pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact, attachment_id
		FROM pariwisata ORDER BY pariwisata_nama ASC
	`

	QueryGetPariwisataByID = `
		SELECT pariwisata_id, pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact, attachment_id
		FROM pariwisata
		WHERE pariwisata_id = $1
	`

	QueryCreatePariwisata = `
		INSERT INTO pariwisata(pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact, attachment_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING pariwisata_id
	`

	QueryUpdatePariwisata = `
		UPDATE pariwisata
		SET pariwisata_nama=$2, pariwisata_lokasi=$3, pariwisata_desc=$4, pariwisata_contact=$5
		WHERE pariwisata_id=$1
	`

	QueryDeletePariwisata = `
		DELETE FROM pariwisata
		WHERE pariwisata_id=$1
	`
	QueryUpdateImagePariwisata = `
		UPDATE pariwisata
		SET attachment_id = $1
		WHERE pariwisata_id = $2
	`
)
