package postgres

const (
	QueryGetAllPariwisata = `
		SELECT pariwisata_id, pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact
		FROM pariwisata
	`

	QueryGetPariwisataByID = `
		SELECT pariwisata_id, pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact
		FROM pariwisata
		WHERE pariwisata_id = $1
	`

	QueryCreatePariwisata = `
		INSERT INTO pariwisata(pariwisata_nama, pariwisata_lokasi, pariwisata_desc, pariwisata_contact)
		VALUES ($1, $2, $3, $4)
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
)
