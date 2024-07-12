package models

type Hotel struct {
	//gorm.Model
	ID          uint    `gorm:"primary_key; not null" json:"id"`
	Name        string  `gorm:"size:255; not null" json:"name"`
	Destination string  `gorm:"size:64; not null" json:"destination"`
	Price       float64 `gorm:"null" json:"price"`
	Image       string  `gorm:"size:255; not null" json:"image"`
	Status      int64   `gorm:"default:0;not null" json:"status"`
	Address     string  `gorm:"size:255;null;" json:"address"`
	Longitude   float64 `gorm:"default:0;not null" json:"longitude"`
	Latitude    float64 `gorm:"default:0;not null" json:"latitude"`
}
