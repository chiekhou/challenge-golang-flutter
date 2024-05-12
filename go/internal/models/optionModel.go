package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Name    string `gorm:"not null" json:"name"`
	Type    string `gorm:"not null" json:"type"`
	HotelID uint   `gorm:"foreignkey: HotelId" json:"hotel_id"`
}
