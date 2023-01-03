package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `gorm:"type:longtext" json:"-"`
}
