package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	UserName string `gorm:"type:varchar(100);unique;not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"not null"`
}
