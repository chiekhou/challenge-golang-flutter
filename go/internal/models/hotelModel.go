package models

import (
	"gorm.io/gorm"
	"time"
)

type Hotel struct {
	gorm.Model
	Checkin   time.Time  `gorm:"not null" json:"checkin"`
	Checkout  time.Time  `gorm:"not null" json:"checkout"`
	Options   []Option   `gorm:"foreignKey:HotelID" json:"options"`
	Feedbacks []Feedback `gorm:"foreignkey:HotelID" json:"feedbacks"`
}
