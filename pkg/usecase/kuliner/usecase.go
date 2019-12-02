package kuliner

type Usecase interface {
	//GetAllResto() ([]models.Restoran, error)
	CreateKuliner(restoID string, name string, desc string, price string, attachmentID int64) error
}
