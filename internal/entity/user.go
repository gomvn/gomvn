package entity

type User struct {
	Id        string `gorm:"primary_key;size:36;not null"`
	TokenHash string `gorm:"primary_key;size:60;not null"`
}
