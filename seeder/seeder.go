package seeder

import (
	"encoding/csv"
	"fmt"
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

func PopulateUserCSV() {
	csvFile, err := os.Create("./db_script/user.csv")
	if err != nil {
		fmt.Println(err)
	}
	csvWriter := csv.NewWriter(csvFile)
	for i:=1; i<1000; i++ {
		err := csvWriter.Write([]string{
			strconv.Itoa(i),
			GenerateField(names),
			RandStringRunes(10)})
		if err != nil {
			log.Fatal(err)
		}
	}
	csvWriter.Flush()
	err = csvFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func PopulateProductCSV() {
	csvFile, err := os.Create("./db_script/product.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(csvFile)
	for i:=1; i<1000; i++ {
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
	csvWriter.Flush()
	err = csvFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
