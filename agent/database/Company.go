package database

import (
	"github.com/XML/agent/model"
)

func GetCompanies() []model.Company {
	db := DBConn
	var comps []model.Company
	db.Find(&comps, "active = ?", true)
	return comps
}

func GetCompany(id string) model.Company {
	db := DBConn
	var comp model.Company
	db.Find(&comp, "id = ? AND active = ?", id, true)
	LoadPosts(id, &comp)
	return comp
}

func CreateCompanies(comp *model.Company) error {
	db := DBConn
	comp.Active = false
	result := db.Create(&comp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
