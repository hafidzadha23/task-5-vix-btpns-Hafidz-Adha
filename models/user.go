package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email" validate:"email"`
	Password string `gorm:"not null;min:6" json:"password" validate:"min=6"`
	Photos   []Photo
}
