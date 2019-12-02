package postgres

const (
	QueryInsertKuliner = `
		INSERT INTO kuliner (
			resto_id,
			kuliner_nama,
			kuliner_harga,
			kuliner_desc,
			attachment_id
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`
)
