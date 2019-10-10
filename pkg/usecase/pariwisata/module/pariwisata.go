package module

import (
	"github.com/TobaTourism/pkg/models"
)

func (u *pariwisata) GetAllPariwisata() ([]models.Pariwisata, error) {
	allPariwisata, err := u.pariwisataRepo.GetAllPariwisata()
	// if err != nil {
	// 	log.Println(err)
	// }

	return allPariwisata, err
}
