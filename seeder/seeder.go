package seeder

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var names = []string {"Lisa","Jennie","Rose","Jisoo",
	"Yeri","Irene","Seulgi","Wendy","Joy",
	"Felix","Bang Chan", "Hyunjin", "Jeongin",
	"Jisung", "Changbin", "Minho", "Seungmin"}

var products = []string {"Cactus","Laptop","IPhone","Tarot cards",
	"Candle","Blazer","Coffee","Plushie","Kettle",
	"Sunscreen", "Meds", "Juice", "Ice",
	"Beer", "Candies", "TV", "Rice"}

func GenerateField(field []string ) string {
	rand.Seed(time.Now().UnixNano())
	name := field[rand.Intn(len(field)-1)]
	return name
}

func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func PopulateCSV(name string) {
	var filepath string
	switch name {
	case "user":
		filepath = "./db_script/user.csv"
	case "product":
		filepath = "./db_script/product.csv"
	}
	csvFile, err := os.Create(filepath)
	csvWriter := csv.NewWriter(csvFile)
	for i:=1; i<=1000; i++ {
		switch name {
		case "user":
			err := csvWriter.Write([]string{
				strconv.Itoa(i),
				GenerateField(names),
				RandStringRunes(10)})
			if err != nil {
				log.Fatal(err)
			}
		case "product":
			err := csvWriter.Write([]string{
				strconv.Itoa(i),
				GenerateField(products),
				strconv.Itoa(rand.Intn(500)),
				strconv.Itoa(rand.Intn(1000)),
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	csvWriter.Flush()
	err = csvFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
