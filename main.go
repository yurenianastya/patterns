package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"patterns/labmvc/controller"
	"patterns/labmvc/seeder"
	_ "patterns/labmvc/seeder"
)

func main() {
	dsn := "root:mypass@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{

	})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	seeder.PopulateUserCSV()
	seeder.PopulateProductCSV()
	controller.RetrieveUsers(database)
	controller.RetrieveProducts(database)
}