package model

type User struct {
	Id int  `gorm:"primaryKey;autoIncrement;"`
	Username string
	Password string `gorm:"foreignKey:UserNumber;references:Id"`

}
