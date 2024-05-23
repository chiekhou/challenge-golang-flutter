package models

type Option struct {
	//gorm.Model
	ID      uint   `gorm:"primary_key; not null" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	Type    string `gorm:"not null" json:"type"`
	HotelID uint   `gorm:"foreignkey: HotelId" json:"hotel_id"`
}
