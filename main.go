package main

import (
	"github.com/lahyri/imobiliaria/app/housing"
	"github.com/lahyri/imobiliaria/config"

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// dsn := viper.GetString(`database`)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// config := config.Config{db}
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	config := config.Config{}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/states", housing.GetStates(&config))
	e.GET("/cities/:state", housing.GetCities(&config))
	e.GET("/districts/:city/:state", housing.GetDistricts(&config))
	e.GET("/", housing.GetAll(&config))
	e.GET("/:uuid", housing.GetOne(&config))
	e.POST("/", housing.Post(&config))
	e.DELETE("/:uuid", housing.Delete(&config))
	e.Logger.Fatal(e.Start(":1323"))
}
