package module

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TobaTourism/pkg/models"
)

// func (u *resto) GetAllResto() ([]models.Restoran, error) {
// 	allResto, err := u.restoRepo.GetAllResto()
// 	if err != nil {
// 		log.Println(err)
// 		return allResto, err
// 	}

// 	return allResto, err
// }

func (u *kuliner) CreateKuliner(restoID string, name string, desc string, price string, attachmentID int64) error {
	var culinary models.Culinary

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID] Error : ", err)
		return err
	}

	culinary.RestoID = restoIDInt
	if name == "" {
		err := fmt.Errorf("[Kuliner][Usecase][name] empty field name")
		log.Println(err)
		return err
	}
	culinary.Name = name

	if desc == "" {
		err := fmt.Errorf("[Kuliner][Usecase][desc] empty field name")
		log.Println(err)
		return err
	}
	culinary.Desc = desc

	priceInt, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse Price] Error : ", err)
		return nil
	}
	culinary.Price = priceInt

	if attachmentID == 0 {
		err := fmt.Errorf("[Kuliner][Usecase][AttachmentID] empty field AttachmentID")
		log.Println(err)
		return err
	}
	culinary.AttachmentID = attachmentID

	err = u.kulinerRepo.CreateKuliner(culinary)
	if err != nil {
		log.Println("[Kuliner][Usecase][CreateKuliner] Error : ", err)
		return err
	}

	return err
}
