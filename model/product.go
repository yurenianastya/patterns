package model

type Product struct {
	Id  int   `gorm:"primaryKey;autoIncrement;"`
	Price      int
	Name       string
	UserNumber int
}


