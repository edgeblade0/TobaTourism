package module

import (
	"fmt"
	"log"
	"strconv"

	"github.com/if416021/TobaTourism/pkg/models"
)

func (u *kuliner) GetAllCulinary(restoID int64) ([]models.Culinary, error) {
	allCulinary, attachmentID, err := u.kulinerRepo.GetAllKuliner(restoID)
	if err != nil {
		log.Println("[Kuliner][Usecase][GetAllKuliner] Error : ", err)
		return allCulinary, err
	}

	for i, id := range attachmentID {
		allCulinary[i].Attachment, err = u.attachmentRepo.GetAttachment(id)
		if err != nil {
			log.Println("[Kuliner][Usecase][GetAttachment] Error : ", err)
			return allCulinary, err
		}

	}
	return allCulinary, err
}

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

	// if desc == "" {
	// 	err := fmt.Errorf("[Kuliner][Usecase][desc] empty field Desc")
	// 	log.Println(err)
	// 	return err
	// }
	culinary.Desc = desc

	priceInt, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse Price] Error : ", err)
		return err
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

func (u *kuliner) GetDetailKuliner(restoID string, culinaryID string) (models.Culinary, error) {
	var culinary models.Culinary

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID getByID ] Error : ", err)
		return culinary, err
	}

	culinaryIDInt, err := strconv.ParseInt(culinaryID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse culinaryID getByID ] Error : ", err)
		return culinary, err
	}

	culinary, attachmentID, err := u.kulinerRepo.GetDetailKuliner(restoIDInt, culinaryIDInt)
	if err != nil {
		log.Println("[Kuliner][Usecase][GetDetailKuliner] Error : ", err)
		return culinary, err
	}

	culinary.Attachment, err = u.attachmentRepo.GetAttachment(attachmentID)
	if err != nil {
		log.Println("[Kuliner][Usecase][GetAttachment] Error : ", err)
		return culinary, err
	}

	return culinary, err
}

func (k *kuliner) UpdateKuliner(restoID string, kulinerID string, kulinerName string, kulinerDesc string, kulinerPrice string) error {
	var culinary models.Culinary

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Update Culinary] Error : ", err)
		return err
	}
	culinary.RestoID = restoIDInt

	kulinerIDInt, err := strconv.ParseInt(kulinerID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Update Culinary] Error : ", err)
		return err
	}
	culinary.CulinaryID = kulinerIDInt

	if kulinerName == "" {
		err = fmt.Errorf("[Usecase][Kuliner][kulinerName Update] no name")
		log.Println(err)
		return err
	}
	culinary.Name = kulinerName

	if kulinerDesc == "" {
		err = fmt.Errorf("[Usecase][Kuliner][KulinerDesc Update] no desc")
		log.Println(err)
		return err
	}
	culinary.Desc = kulinerDesc

	price, err := strconv.ParseInt(kulinerPrice, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse Price on Update] Error : ", err)
		return err
	}
	culinary.Price = price

	err = k.kulinerRepo.UpdateKuliner(culinary)
	if err != nil {
		log.Println("[Kuliner][Usecase][UpdateKuliner] Error : ", err)
		return err
	}

	return err
}

func (k *kuliner) UpdateImageKuliner(restoID string, kulinerID string, attachmentID int64) error {
	var culinary models.Culinary

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Update Culinary] Error : ", err)
		return err
	}
	culinary.RestoID = restoIDInt

	kulinerIDInt, err := strconv.ParseInt(kulinerID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Update Culinary] Error : ", err)
		return err
	}
	culinary.CulinaryID = kulinerIDInt

	if attachmentID == 0 {
		err = fmt.Errorf("[Usecase][Kuliner][attachmentID on Update] no attachmentID")
		log.Println(err)
		return err
	}
	culinary.AttachmentID = attachmentID

	err = k.kulinerRepo.UpdateImageKuliner(culinary)
	if err != nil {
		log.Println("[Restoran][Usecase][UpdateKuliner] Error : ", err)
		return err
	}

	return err
}

func (k *kuliner) DeleteKuliner(restoID string, kulinerID string) error {
	var culinary models.Culinary

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Delete Culinary] Error : ", err)
		return err
	}
	culinary.RestoID = restoIDInt

	kulinerIDInt, err := strconv.ParseInt(kulinerID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Kuliner][Parse RestoID on Delete Culinary] Error : ", err)
		return err
	}
	culinary.CulinaryID = kulinerIDInt

	err = k.kulinerRepo.DeleteKuliner(culinary)
	if err != nil {
		log.Println("[Kuliner][Usecase][UpdateKuliner] Error : ", err)
		return err
	}

	return err
}
