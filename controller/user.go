package controller

import (
	"fmt"
	"gorm.io/gorm"
	"patterns/labmvc/model"
)

var user model.User
var users []model.User

// CreateUser godoc
// @Summary create a new user
// @Description POST method which creates new user object
// @Param user body model.User true "User data"
// @Success 200 {object} model.User
// @Router /user [post]
func CreateUser(db *gorm.DB, username string, password string)  {
	user = model.User{Username: username, Password: password}
	db.Create(&user)
	fmt.Println("User was successfully created")
}

// RetrieveUser godoc
// @Summary get a user
// @Description GET method which returns user object
// @Param id path integer true "User ID"
// @Success 200 {object} model.User
// @Router /user/{id} [get]
func RetrieveUser(db *gorm.DB, id int) {
	db.First(&user, id)
	fmt.Println("User: ", user)
}


// RetrieveUsers godoc
// @Summary get users
// @Description GET method which returns users objects
// @Success 200 {object} model.User
// @Router /users [get]
func RetrieveUsers(db *gorm.DB) {
	db.Find(&users)
	fmt.Println("Users: ", users)
}

// UpdateUser godoc
// @Summary update user
// @Description PUT method which updates user object
// @Param id path integer true "User ID"
// @Success 200 {object} model.User
// @Router /user/{id} [put]
func UpdateUser(db *gorm.DB, id int, username string, password string) {
	newUser := model.User{Id: id, Username: username, Password: password}
	user.Id = id
	db.Model(&user).Updates(newUser)
	fmt.Println("User updated: ", newUser)
}

// DeleteUser godoc
// @Summary delete user
// @Description DELETE method which deletes user object
// @Param id path integer true "User ID"
// @Success 200 {object} model.User
// @Router /user/{id} [delete]
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

