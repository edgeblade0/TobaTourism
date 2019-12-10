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
	QuerySelectAllCulinary = `
		SELECT kuliner_id,kuliner_nama,kuliner_harga,kuliner_desc,attachment_id
		FROM kuliner
		WHERE resto_id = $1
	`

	QuerySelectDetailCulinary = `
		SELECT kuliner_nama,kuliner_harga,kuliner_desc,attachment_id
		FROM kuliner
		WHERE resto_id = $1 AND kuliner_id = $2
	`
	QueryUpdateCulinary = `
		UPDATE kuliner
		SET 
			kuliner_nama = $1,
			kuliner_harga = $2,
			kuliner_desc = $3
		WHERE resto_id = $4 AND kuliner_id = $5
	`

	QueryUpdateImageCulinary = `
		UPDATE kuliner
		SET 
			attachment_id = $1
		WHERE resto_id = $2 AND kuliner_id = $3
	`

	QueryDeleteCulinary = `
		DELETE FROM kuliner
		WHERE resto_id = $1 AND kuliner_id = $2
	`

	QueryDeleteAllCulinary = `
		DELETE FROM kuliner
		WHERE resto_id = $1
	`
)
