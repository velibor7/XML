package database

import "github.com/XML/agent/model"

func GetAdminCompany(id string) model.Company {
	db := DBConn
	var comp model.Company

	db.Find(&comp, "id = ?", id)

	LoadPosts(id, &comp)
	return comp
}

func GetAdminCompanies() []model.Company {
	db := DBConn
	var comps []model.Company
	db.Find(&comps)
	return comps
}

func ActivateCompany(id string) error {
	db := DBConn
	var comp model.Company

	db.Find(&comp, "id = ?", id)
	comp.Active = true
	result := db.Save(&comp)
	AddCompanyToUser(comp.UserName, &comp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddCompanyToUser(username string, comp *model.Company) {
	db := DBConn
	var user model.User
	db.Find(&user, "username = ?", username)
	user.Company = append(user.Company, *comp)
	user.Role = 1
	db.Save(&user)
}
