package models

type Login struct {
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null;min:6" json:"password"`
}
