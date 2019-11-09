package postgres

var queryGetAllPariwisata = "SELECT pariwisata_id, pariwisata_name, pariwisata_lokasi FROM pariwisata"

var queryCreatePariwisata = `INSERT INTO pariwisata (pariwisata_name, pariwisata_lokasi)VALUES($1, $2) RETURNING pariwisata_id`
