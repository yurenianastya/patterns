package model

import (
	"github.com/spf13/viper"
	"log"
)

type Product struct {
	Id  int   `gorm:"primaryKey;autoIncrement;"`
	Price      int
	Name       string
	UserNumber int
}

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


