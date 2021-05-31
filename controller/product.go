package controller

import (
	"fmt"
	"gorm.io/gorm"
	"patterns/labmvc/model"
)

var product model.Product
var products []model.Product

// CreateProduct godoc
// @Summary create a new product
// @Description POST method which creates new product object
// @Param product body model.Product true "Product data"
// @Success 200 {object} model.Product
// @Router /product [post]
func CreateProduct(db *gorm.DB, name string, price int, usernum int)  {
	product = model.Product{Name: name, Price: price, UserNumber: usernum}
	db.Create(&product)
	fmt.Println("Product was successfully created")
}

// RetrieveProduct godoc
// @Summary get a product
// @Description GET method which returns product object
// @Param id path integer true "Product ID"
// @Success 200 {object} model.Product
// @Router /product/{id} [get]
func RetrieveProduct(db *gorm.DB, id int) {
	db.First(&product, id)
	fmt.Println("Product: ", product)
}

// RetrieveProducts godoc
// @Summary get products
// @Description GET method which returns products objects
// @Success 200 {object} model.Product
// @Router /products [get]
func RetrieveProducts(db *gorm.DB) {
	db.Find(&products)
	fmt.Println("Products: ", products)
}

// UpdateProduct godoc
// @Summary update product
// @Description PUT method which updates product object
// @Param id path integer true "Product ID"
// @Success 200 {object} model.Product
// @Router /product/{id} [put]
func UpdateProduct(db *gorm.DB, id int, name string, price int, usernum int) {
	newProduct := model.Product{Id: id, Name: name, Price: price, UserNumber: usernum}
	product.Id = id
	db.Model(&product).Updates(newProduct)
	fmt.Println("User updated: ", newProduct)
}

// DeleteProduct godoc
// @Summary delete product
// @Description DELETE method which deletes product object
// @Param id path integer true "Product ID"
// @Success 200 {object} model.Product
// @Router /product/{id} [delete]
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
