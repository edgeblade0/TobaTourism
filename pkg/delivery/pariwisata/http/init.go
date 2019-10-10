package http

import (
	"github.com/labstack/echo"

	pariwisataUsecase "github.com/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	pariwisataUsecase pariwisataUsecase.Usecase
}

func InitPariwisataHandler(e *echo.Echo, p pariwisataUsecase.Usecase) {
	handler := &pariwisata{
		pariwisataUsecase: p,
	}

	//register handler
	e.GET("/pariwisata", handler.GetAllPariwisata)
}

// ResponseError represent the reseponse error struct
// type ResponseError struct {
// 	Message string `json:"message"`
// }

// ArticleHandler  represent the httphandler for article
// type ArticleHandler struct {
// 	test string
// }

// NewArticleHandler will initialize the articles/ resources endpoint
// func NewArticleHandler(e *echo.Echo) {
// 	handler := &ArticleHandler{
// 		test: "test",
// 	}
// 	e.GET("/articles", handler.FetchArticle)
// }

// FetchArticle will fetch the article based on given params
// func (a *ArticleHandler) FetchArticle(c echo.Context) error {
// 	Conf := config.InitConfig()
// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}
// 	listPar := models.Pariwisata{1, "test", "test"}
// 	log.Println("Tes2t", Conf.Db.Conn)
// 	c.Response().Header().Set(`X-Cursor`, "header")
// 	return c.JSON(http.StatusOK, listPar)
// }
