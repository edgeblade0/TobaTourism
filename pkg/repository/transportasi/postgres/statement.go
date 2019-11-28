package postgres

const (
	QueryGetAllTransportasi = `
		SELECT id, nama, rute, description, contact, harga
		FROM transportasi
	`

	QueryGetTransportasiByID = `
		SELECT id, nama, rute, description, contact, harga
		FROM transportasi
		WHERE id = $1
	`

	QueryCreateTransportasi = `
		INSERT INTO transportasi(nama, rute, description, contact, harga)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	QueryUpdateTransportasi = `
		UPDATE transportasi
		SET nama=$2, rute=$3, description=$4, contact=$5, harga=$6
		WHERE id=$1
	`

	QueryDeleteTransportasi = `
		DELETE FROM transportasi
		WHERE id=$1
	`
)
