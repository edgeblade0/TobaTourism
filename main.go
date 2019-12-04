package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/TobaTourism/middleware"
	"github.com/TobaTourism/pkg/common/config"
	pariwisataDeliver "github.com/TobaTourism/pkg/delivery/pariwisata/http"
	"github.com/TobaTourism/pkg/models"
	pariwisataRepo "github.com/TobaTourism/pkg/repository/pariwisata/postgres"
	pariwisataUseCase "github.com/TobaTourism/pkg/usecase/pariwisata/module"

	restoDeliver "github.com/TobaTourism/pkg/delivery/resto/http"
	restoRepo "github.com/TobaTourism/pkg/repository/resto/postgres"
	restoUseCase "github.com/TobaTourism/pkg/usecase/resto/module"

	attachmentDeliver "github.com/TobaTourism/pkg/delivery/attachment/http"
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment/postgres"
	attachmentUseCase "github.com/TobaTourism/pkg/usecase/attachment/module"

	kulinerDeliver "github.com/TobaTourism/pkg/delivery/kuliner/http"
	kulinerRepo "github.com/TobaTourism/pkg/repository/kuliner/postgres"
	kulinerUseCase "github.com/TobaTourism/pkg/usecase/kuliner/module"
)

var Conf *models.Config

func main() {
	Conf = config.InitConfig()
	//http
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	//DB
	// db := conn.InitDB(Conf.Db.Conn)
	db, err := sql.Open("postgres", Conf.Db.Conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//Konten Pariwisata
	pariwisata(e, db)

	restoran(e, db)
	attachment(e, db)
	kuliner(e, db)

	log.Fatal(e.Start(":9090"))
}

func pariwisata(e *echo.Echo, db *sql.DB) {
	pariwisataRepo := pariwisataRepo.InitPariwisataRepo(db)
	pariwisataUsecase := pariwisataUseCase.InitPariwisataUsecase(pariwisataRepo)
	pariwisataDeliver.InitPariwisataHandler(e, pariwisataUsecase)

}

func restoran(e *echo.Echo, db *sql.DB) {
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)

	kulinerRepo := kulinerRepo.InitKulinerRepo(db)
	kulinerUsecase := kulinerUseCase.InitKulinerUsecase(kulinerRepo, attachmentRepo)

	restoRepo := restoRepo.InitRestoRepo(db)
	restoUsecase := restoUseCase.InitRestoUsecase(restoRepo, attachmentRepo, kulinerRepo, kulinerUsecase)
	restoDeliver.InitRestoHandler(e, restoUsecase, attachmentUseCase)
}

func attachment(e *echo.Echo, db *sql.DB) {
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)
	attachmentDeliver.InitAttachmentHandler(e, attachmentUseCase)
}

func kuliner(e *echo.Echo, db *sql.DB) {
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)

	kulinerRepo := kulinerRepo.InitKulinerRepo(db)
	kulinerUsecase := kulinerUseCase.InitKulinerUsecase(kulinerRepo, attachmentRepo)
	kulinerDeliver.InitKulinerHandler(e, kulinerUsecase, attachmentUseCase)
}
