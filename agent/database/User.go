package database

import "github.com/XML/agent/model"

func GetUsers() []model.User {
	db := DBConn
	var users []model.User
	db.Find(&users)
	return users
}

func RegisterUsers(user *model.User) error {
	db := DBConn
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetUser(username string) *model.User {
	db := DBConn
	oldUser := new(model.User)
	db.Find(&oldUser, "username = ?", username)
	return oldUser
}
