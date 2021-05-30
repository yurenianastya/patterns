package controller

import (
	"fmt"
	"gorm.io/gorm"
	"patterns/labmvc/model"
)

var user model.User
var users []model.User

func CreateUser(db *gorm.DB, username string, password string)  {
	user = model.User{Username: username, Password: password}
	db.Create(&user)
	fmt.Println("User was successfully created")
}
func RetrieveUser(db *gorm.DB, id int) {
	db.First(&user, id)
	fmt.Println("User: ", user)
}

func RetrieveUsers(db *gorm.DB) {
	db.Find(&users)
	fmt.Println("Users: ", users)
}

func UpdateUser(db *gorm.DB, id int, username string, password string) {
	newUser := model.User{Id: id, Username: username, Password: password}
	user.Id = id
	db.Model(&user).Updates(newUser)
	fmt.Println("User updated: ", newUser)
}

func DeleteUser(db *gorm.DB, UserId int)  {
	var index int
	for p, v := range users {
		if v.Id == UserId {
			index = p
		}
	}
	db.Delete(&users, UserId)
	if len(users) > 1 {
		users = append(users[:index], users[index:]...)
	} else {
		users = nil
	}
	fmt.Println("User deleted with this id: ", UserId)
}

