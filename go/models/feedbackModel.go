package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Note        int    `gorm:"note:int"`
	Commentaire string `gorm:"commentaire:varchar(30);"`
	UserID      uint   `gorm:"foreignkey:UserID"`
	HotelID     uint   `gorm:"foreignkey:HotelID"`
}
