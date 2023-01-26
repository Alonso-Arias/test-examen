package main

import (
	"context"
	"net/http"

	security "github.com/alonso/test/context"
	errs "github.com/alonso/test/errors"
	"github.com/alonso/test/log"
	"github.com/alonso/test/services/local"
	"github.com/alonso/test/services/model"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/alonso/test/api/docs"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

var basicData = model.GetLocales{}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()

	e.GET("/api/v1/pharmacy/locals/:commune", getLocalsByCommune)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// al iniciar se crea una goroutine
	go getData()
	e.Logger.Fatal(e.Start(":1323"))

}

// en caso de extraer datos para ser consultados se realizo
// un metodo asyncronico para rescatar los datos sin que se quede colgado
// en el caso de haber muchos
func getData() {

	log := loggerf.WithField("func", "getData")

	log.Debugf("Init")
	defer log.Debugf("End")

	ls, err := security.GetLocals()
	if err != nil {
		log.WithError(err).Error("problems with getting data")
	}

	basicData = ls

}

// get locals by commune godoc
// @Summary Get locals
// @tags Locals
// @Description obtiene locales al ingresarse una comuna
// @ID getLocalsByCommune
// @Accept  json
// @Produce  json
// @Param commune path string true "Comuna"
// @Success 200  {object} local.GetLocalsByComunaResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /pharmacy/locals/{commune} [get]
func getLocalsByCommune(c echo.Context) error {
	// se arma el request con ya la data cargada
	// si no se carga retornara un error
	if len(basicData) == 0 {
		return c.JSON(http.StatusInternalServerError, "no data loaded")
	}
	req := local.GetLocalsByComunaRequest{
		Data:    basicData,
		Commune: c.Param("commune"),
	}

	res, err := local.LocalService{}.GetLocalsByCommune(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
