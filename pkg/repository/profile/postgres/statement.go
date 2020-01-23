package postgres

const (
	QueryGetProfile = `
		SELECT name, email, no_hp, alamat, dob, social_media
		FROM profiles
		WHERE account_id = $1
	`
)
