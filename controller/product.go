package controller

import (
	"fmt"
	"gorm.io/gorm"
	"patterns/labmvc/model"
)

var product model.Product
var products []model.Product

func CreateProduct(db *gorm.DB, name string, price int, usernum int)  {
	product = model.Product{Name: name, Price: price, UserNumber: usernum}
	db.Create(&product)
	fmt.Println("Product was successfully created")
}

func RetrieveProduct(db *gorm.DB, id int) {
	db.First(&product, id)
	fmt.Println("Product: ", product)
}

func RetrieveProducts(db *gorm.DB) {
	db.Find(&products)
	fmt.Println("Products: ", products)
}

func UpdateProduct(db *gorm.DB, id int, name string, price int, usernum int) {
	newProduct := model.Product{Id: id, Name: name, Price: price, UserNumber: usernum}
	product.Id = id
	db.Model(&user).Updates(newProduct)
	fmt.Println("User updated: ", newProduct)
}

func DeleteProduct(db *gorm.DB, id int)  {
	var index int
	for p, v := range products {
		if v.Id == id {
			index = p
		}
	}
	db.Delete(&products, id)
	if len(users) > 1 {
		products = append(products[:index], products[index:]...)
	} else {
		products = nil
	}
	fmt.Println("Product deleted with this id: ", id)
}
