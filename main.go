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

	log.Fatal(e.Start(":9090"))
}

func pariwisata(e *echo.Echo, db *sql.DB) {
	pariwisataRepo := pariwisataRepo.InitPariwisataRepo(db)
	pariwisataUsecase := pariwisataUseCase.InitPariwisataUsecase(pariwisataRepo)
	pariwisataDeliver.InitPariwisataHandler(e, pariwisataUsecase)

}
