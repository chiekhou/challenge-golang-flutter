package models

import (
	"gorm.io/gorm"
	"time"
)

type Hotel struct {
	gorm.Model
	Checkin   time.Time `gorm:"type:varchar(30);not null"`
	Checkout  time.Time `gorm:"type:varchar(30);not null"`
	Options   string
	Feedbacks []Feedback `gorm:"foreignkey:HotelID"`
}
