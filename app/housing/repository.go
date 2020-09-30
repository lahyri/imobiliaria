package housing

import (
	"errors"

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

func getAllMock() ([]Housing, error) {

	return Database, nil
}

func saveMock(housing Housing) (Housing, error) {

	uuid, err := helper.GenerateUUID()
	if err != nil {
		return Housing{}, err
	}
	housing.UUID = uuid
	Database = append(Database, housing)
	return housing, nil

}

func getOneMock(uuid string) (Housing, error) {
	for _, housing := range Database {
		if housing.UUID == uuid {
			return housing, nil
		}
	}
	return Housing{}, errors.New("Housing not found")

}

func deleteMock(uuid string) error {
	for i, housing := range Database {
		if housing.UUID == uuid {
			Database = append(Database[:i], Database[i+1:]...)
			return nil
		}
	}
	return errors.New("Housing not found")

}

func getStatesMock() ([]string, error) {

	states := make(map[string]bool)
	for _, housing := range Database {
		_, ok := states[housing.State]
		if !ok {
			states[housing.State] = true
		}
	}

	var response []string
	for state, _ := range states {
		response = append(response, state)
	}
	return response, nil
}

func getCitiesMock(state string) ([]string, error) {

	cities := make(map[string]bool)
	for _, housing := range Database {
		if housing.State == state {
			_, ok := cities[housing.State]
			if !ok {
				cities[housing.State] = true
			}
		}

	}

	var response []string
	for city, _ := range cities {
		response = append(response, city)
	}
	return response, nil
}

func getDistrictsMock(state string, city string) ([]string, error) {

	districts := make(map[string]bool)
	for _, housing := range Database {
		if housing.State == state {
			if housing.City == city {
				_, ok := districts[housing.State]
				if !ok {
					districts[housing.State] = true
				}
			}

		}

	}

	var response []string
	for district, _ := range districts {
		response = append(response, district)
	}
	return response, nil
}
