package helper

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Filter struct {
	Ints    map[string]int
	Strings map[string]string
	Bools   map[string]bool
	Floats  map[string]float64
}

var (
	stringFilter = []string{
		"description",
		"street",
		"district",
		"city",
		"state",
		"cep",
	}

	intFilter = []string{
		"type",
		"room_amount",
		"suite_amount",
		"bedroom_amount",
		"parking_spaces",
		"area",
		"number",
		"floor",
	}

	boolFilter = []string{
		"has_cabinets",
		"has_concierge",
	}

	floatFilter = []string{
		"fees",
		"value",
	}
)

func BuildFilter(c echo.Context) Filter {
	stringMap := make(map[string]string)
	for _, key := range stringFilter {
		val := c.QueryParam(key)
		if val != "" {
			stringMap[key] = val
		}
	}

	intMap := make(map[string]int)
	for _, key := range intFilter {
		val := c.QueryParam(key)
		if val != "" {
			intMap[key], _ = strconv.Atoi(val)
		}
	}

	boolMap := make(map[string]bool)
	for _, key := range boolFilter {
		val := c.QueryParam(key)
		if val != "" {
			boolMap[key], _ = strconv.ParseBool(val)
		}
	}

	floatMap := make(map[string]float64)
	for _, key := range floatFilter {
		val := c.QueryParam(key)
		if val != "" {
			floatMap[key], _ = strconv.ParseFloat(val, 32)
		}
	}

	return Filter{
		Ints:    intMap,
		Strings: stringMap,
		Bools:   boolMap,
		Floats:  floatMap,
	}
}

func BuildFilters(db *gorm.DB, filter Filter) *gorm.DB {
	if len(filter.Strings) != 0 {
		for key, val := range filter.Strings {
			db.Where(key+" = ?", val)
		}
	}
	if len(filter.Ints) != 0 {
		for key, val := range filter.Ints {
			db.Where(key+" = ?", val)
		}
	}
	if len(filter.Bools) != 0 {
		for key, val := range filter.Bools {
			db.Where(key+" = ?", val)
		}
	}
	if len(filter.Floats) != 0 {
		for key, val := range filter.Floats {
			db.Where(key+" = ?", val)
		}
	}
	return db
}
