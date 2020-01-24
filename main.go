package main

import (
	sql "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/TobaTourism/middleware"
	"github.com/TobaTourism/pkg/common/config"
	attachmentDeliver "github.com/TobaTourism/pkg/delivery/attachment/http"
	experienceDeliver "github.com/TobaTourism/pkg/delivery/experience/http"
	kulinerDeliver "github.com/TobaTourism/pkg/delivery/kuliner/http"
	pariwisataDeliver "github.com/TobaTourism/pkg/delivery/pariwisata/http"
	profileDeliver "github.com/TobaTourism/pkg/delivery/profile/http"
	restoDeliver "github.com/TobaTourism/pkg/delivery/resto/http"
	transportasiDeliver "github.com/TobaTourism/pkg/delivery/transportasi/http"
	"github.com/TobaTourism/pkg/models"
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment/postgres"
	experienceRepo "github.com/TobaTourism/pkg/repository/experience/postgres"
	kulinerRepo "github.com/TobaTourism/pkg/repository/kuliner/postgres"
	pariwisataRepo "github.com/TobaTourism/pkg/repository/pariwisata/postgres"
	profileRepo "github.com/TobaTourism/pkg/repository/profile/postgres"
	restoRepo "github.com/TobaTourism/pkg/repository/resto/postgres"
	transportasiRepo "github.com/TobaTourism/pkg/repository/transportasi/postgres"
	attachmentUseCase "github.com/TobaTourism/pkg/usecase/attachment/module"
	experienceUseCase "github.com/TobaTourism/pkg/usecase/experience/module"
	kulinerUseCase "github.com/TobaTourism/pkg/usecase/kuliner/module"
	pariwisataUseCase "github.com/TobaTourism/pkg/usecase/pariwisata/module"
	profileUseCase "github.com/TobaTourism/pkg/usecase/profile/module"
	restoUseCase "github.com/TobaTourism/pkg/usecase/resto/module"
	transportasiUseCase "github.com/TobaTourism/pkg/usecase/transportasi/module"
)

var Conf *models.Config

func main() {
	Conf = config.InitConfig()
	//http
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	// e.Use(middL.JwtAuthentication)

	//DB
	// db := conn.InitDB(Conf.Db.Conn)
	db, err := sql.Open("postgres", Conf.Db.Conn)
	dbAuth, err := sql.Open("postgres", Conf.Db.ConnAuth)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Start all services
	startService(e, db)

	restoran(e, db)
	attachment(e, db)
	kuliner(e, db)
	pariwisata(e, db)
	profile(e, dbAuth)

	log.Fatal(e.Start(":9090"))
}

func startService(e *echo.Echo, db *sql.DB) {

	experienceRepo := experienceRepo.InitExperienceRepo(db)
	experienceUsecase := experienceUseCase.InitExperienceUsecase(experienceRepo)
	experienceDeliver.InitExperienceHandler(e, experienceUsecase)

	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)

	transportasiRepo := transportasiRepo.InitTransportasiRepo(db)
	transportasiUsecase := transportasiUseCase.InitTransportasiUsecase(transportasiRepo, attachmentRepo)
	transportasiDeliver.InitTransportasiHandler(e, transportasiUsecase, attachmentUseCase)
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

func pariwisata(e *echo.Echo, db *sql.DB) {
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)
	pariwisataRepo := pariwisataRepo.InitPariwisataRepo(db)
	pariwisataUsecase := pariwisataUseCase.InitPariwisataUsecase(pariwisataRepo, attachmentRepo)
	pariwisataDeliver.InitPariwisataHandler(e, pariwisataUsecase, attachmentUseCase)
}

func profile(e *echo.Echo, db *sql.DB) {
	profileRepo := profileRepo.InitProfileRepo(db)
	profileUsecase := profileUseCase.InitProfileUsecase(profileRepo)
	profileDeliver.InitProfileHandler(e, profileUsecase)
}
