package module

import (

	// "log"
	"github.com/TobaTourism/pkg/models"
)

func (u *pariwisata) GetAllPariwisata() ([]models.Pariwisata, error) {
	allPariwisata, err := u.pariwisataRepo.GetAllPariwisata()
	// if err != nil {
	// 	log.Println(err)
	// }

	return allPariwisata, err
}

//for the insert
// func (u *pariwisata) CreatePariwisata() ([]models.Pariwisata, error) {
// 	allPariwisata, err := u.pariwisataRepo.CreatePariwisata()
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }

// 	// return allPariwisata, err
// }

func (u *pariwisata) CreatePariwisata(pName, pLokasi string) error {
	// allPariwisata, err := u.pariwisataRepo.CreatePariwisata()
	// if err != nil {
	// 	log.Println(err)
	// }
	err := u.pariwisataRepo.CreatePariwisata(pName, pLokasi)

	// log.Println(pName)
	return err

}
