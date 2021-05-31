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


func CreateProduct(w http.ResponseWriter, r *http.Request) {
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
	var product model.Product
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&product)
	fmt.Println("Endpoint Hit: createProduct")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Fatal(err)
	}
}

func RetrieveProducts(w http.ResponseWriter, r *http.Request)  {
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
	var products []model.Product
	db.Find(&products)
	fmt.Println("Endpoint Hit: retrieveProduct")
	json.NewEncoder(w).Encode(products)
}

func RetrieveProduct(w http.ResponseWriter, r *http.Request)  {
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
	var product model.Product
	db.First(&product, key)
	fmt.Println("Endpoint Hit: retrieveProduct")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
	var product model.Product
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		log.Fatal(err)
	}
	product.Id, _ = strconv.Atoi(key)
	db.Model(&product).Updates(product)
	fmt.Println("Endpoint Hit: updateProduct")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
	var product model.Product
	db.Delete(&product, key)
	fmt.Println("Endpoint Hit: deleteProduct")
	json.NewEncoder(w).Encode(product)
}