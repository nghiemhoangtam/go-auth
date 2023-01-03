package models

type User struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password []byte `gorm:"type:longtext"`
}
