package models

type Photo struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `gorm:"not null" json:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	User     User
}
