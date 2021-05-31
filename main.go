package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	_ "patterns/labmvc/docs"
	"patterns/labmvc/view"
)

func GetEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to homepage!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8000/")
	log.Println("Quit the server with CONTROL-C.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/users", view.RetrieveUsers).Methods("GET")
	router.HandleFunc("/user/{id}", view.RetrieveUser).Methods("GET")
	router.HandleFunc("/user", view.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", view.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", view.DeleteUser).Methods("DELETE")
	router.HandleFunc("/products", view.RetrieveProducts).Methods("GET")
	router.HandleFunc("/product/{id}", view.RetrieveProduct).Methods("GET")
	router.HandleFunc("/product", view.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", view.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", view.DeleteProduct).Methods("DELETE")
	err := http.ListenAndServe(GetEnvVariable("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
// @title Amazon Swagger API
// @version 1.0
// @description Swagger API for Golang Project.
// @termsOfService http://swagger.io/terms/
func main() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Run()
	handleRequests()
}