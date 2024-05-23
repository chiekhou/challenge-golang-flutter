package models

import (
	"time"
)

type Hotel struct {
	//gorm.Model
	ID        uint       `gorm:"primary_key; not null" json:"id"`
	Checkin   time.Time  `gorm:"not null" json:"checkin"`
	Checkout  time.Time  `gorm:"not null" json:"checkout"`
	Options   []Option   `gorm:"foreignKey:HotelID" json:"options"`
	Feedbacks []Feedback `gorm:"foreignkey:HotelID" json:"feedbacks"`
}
