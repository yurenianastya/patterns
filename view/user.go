package view

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"patterns/labmvc/model"
	"strconv"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(model.GetEnvVariable("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&user)
	fmt.Println("Endpoint Hit: createUser")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func RetrieveUsers(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open(mysql.Open(model.GetEnvVariable("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var users []model.User
	db.Find(&users)
	fmt.Println("Endpoint Hit: retrieveUsers")
	json.NewEncoder(w).Encode(users)
}

func RetrieveUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open(mysql.Open(model.GetEnvVariable("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	vars := mux.Vars(r)
	key := vars["id"]
	var user model.User
	db.First(&user, key)
	fmt.Println("Endpoint Hit: retrieveUser")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(model.GetEnvVariable("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Fatal(err)
	}
	user.Id, _ = strconv.Atoi(key)
	db.Model(&user).Updates(user)
	fmt.Println("Endpoint Hit: updateUser")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(model.GetEnvVariable("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	vars := mux.Vars(r)
	key := vars["id"]
	var user model.User
	db.Delete(&user, key)
	fmt.Println("Endpoint Hit: deleteUser")
	json.NewEncoder(w).Encode(user)
}