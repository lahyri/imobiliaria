package housing

import (
	"gorm.io/gorm"

	"github.com/lahyri/imobiliaria/helper"
)

func getAll(db *gorm.DB, filter *helper.Filter) ([]Housing, error) {

	var housings []Housing
	db = helper.BuildFilters(db, *filter)
	resp := db.Find(&housings)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return housings, nil
}

func save(db *gorm.DB, housing Housing) (Housing, error) {
	var err error
	housing.UUID, err = helper.GenerateUUID()
	if err != nil {
		return Housing{}, err
	}
	resp := db.Create(&housing)
	if resp.Error != nil {
		return Housing{}, resp.Error
	}
	return getOne(db, housing.UUID)

}

func getOne(db *gorm.DB, uuid string) (Housing, error) {
	housing := Housing{}
	resp := db.First(&housing, "uuid = ?", uuid)
	if resp.Error != nil {
		return Housing{}, resp.Error
	}
	return housing, nil

}

func delete(db *gorm.DB, uuid string) error {
	resp := db.Delete(&Housing{}, "uuid = ?", uuid)
	if resp.Error != nil {
		return resp.Error
	}
	return nil

}

func getStates(db *gorm.DB) ([]string, error) {

	var states []string
	resp := db.Model(&Housing{}).Distinct().Pluck("state", &states)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return states, nil
}

func getCities(db *gorm.DB, state string) ([]string, error) {

	var cities []string
	resp := db.Model(&Housing{}).Distinct().Pluck("state", &cities).Where("state = ?", state)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return cities, nil
}

func getDistricts(db *gorm.DB, state string, city string) ([]string, error) {

	var districts []string
	resp := db.Model(&Housing{}).Distinct().Pluck("state", &districts).Where("state, city = ?, ?", state, city)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return districts, nil
}
