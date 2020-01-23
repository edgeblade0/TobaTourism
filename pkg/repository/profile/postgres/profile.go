package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *profile) GetProfile(userID int64) (models.Profile, error) {
	profile := models.Profile{}

	statement, err := r.DB.Prepare(QueryGetProfile)
	if err != nil {
		log.Println("[Repository][Profile][Prepare Get] Error : ", err)
		return profile, err
	}

	defer statement.Close()

	rows, err := statement.Query(userID)
	if err != nil {
		log.Println("[Repository][Profile][Query Get] Error : ", err)
		return profile, err
	}

	for rows.Next() {
		err := rows.Scan(&profile.Name, &profile.Email, &profile.NoHP, &profile.Alamat, &profile.DOB, &profile.SocialMedia)
		if err != nil {
			log.Println("[Repository][Profile][Row Next Get] Error : ", err)
			return profile, err
		}
	}

	return profile, nil
}

func (r *profile) UpdateProfile(profile models.Profile) error {

	return nil
}
