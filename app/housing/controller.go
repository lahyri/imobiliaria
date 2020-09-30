package housing

import (
	"net/http"

	"github.com/lahyri/imobiliaria/config"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ReqHousing struct {
	Type          int      `json:"type"`
	RoomAmount    int      `json:"room_amount"`
	SuiteAmount   int      `json:"suite_amount"`
	BedroomAmount int      `json:"bedroom_amount"`
	ParkingSpaces int      `json:"parking_spaces"`
	Area          int      `json:"area"`
	HasCabinets   bool     `json:"has_cabinets"`
	Description   string   `json:"description"`
	Fees          *float64 `json:"fees"`
	HasConcierge  *bool    `json:"has_concierge"`
	Street        string   `json:"street"`
	Number        int      `json:"number"`
	District      string   `json:"district"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	CEP           string   `json:"cep"`
	Floor         *int     `json:"floor"`
	RentValue     float64  `json:"rent_value"`
}

type RespHousing struct {
	UUID          string   `json:"uuid"`
	Type          int      `json:"type"`
	RoomAmount    int      `json:"room_amount"`
	SuiteAmount   int      `json:"suite_amount"`
	BedroomAmount int      `json:"bedroom_amount"`
	ParkingSpaces int      `json:"parking_spaces"`
	Area          int      `json:"area"`
	HasCabinets   bool     `json:"has_cabinets"`
	Description   string   `json:"description"`
	Fees          *float64 `json:"fees,omitempty"`
	HasConcierge  *bool    `json:"has_concierge,omitempty"`
	Street        string   `json:"street"`
	Number        int      `json:"number"`
	District      string   `json:"district"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	CEP           string   `json:"cep"`
	Floor         *int     `json:"floor,omitempty"`
	RentValue     float64  `json:"rent_value"`
}

func GetDistricts(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		state := c.Param("state")
		city := c.Param("city")

		//districts, err := getDistricts(config.DB, state, city)
		districts, err := getDistrictsMock(state, city)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, districts)
	}
}
func GetCities(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		state := c.Param("state")

		// cities, err := getCities(config.DB, state)
		cities, err := getCitiesMock(state)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, cities)
	}
}
func GetStates(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		// states, err := getStates(config.DB)
		states, err := getStatesMock()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, states)
	}
}

func GetAll(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		response, err := getAllMock()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, response)
	}
}

func GetOne(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		uuid := c.Param("uuid")
		response, err := getOneMock(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, response)
	}
}
func Delete(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		uuid := c.Param("uuid")
		err := deleteMock(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
func Post(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := ReqHousing{}

		c.Bind(&req)
		// defer c.Request().Body.Close()
		// err := json.NewDecoder(c.Request().Body).Decode(&req)
		// if err != nil {
		// 	log.Fatalf("Failed reading the request body %s", err)
		// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		// }
		house := Housing{}

		copier.Copy(&house, req)

		housing, err := saveMock(house)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, housing)
	}
}
