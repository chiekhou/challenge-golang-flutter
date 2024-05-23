package models

type Feedback struct {
	//gorm.Model
	ID          uint   `gorm:"primary_key; not null" json:"id"`
	Note        int    `gorm:"not null" json:"note"`
	Commentaire string `gorm:"size:30; not null" json:"commentaire"`
	UserID      uint   `gorm:"foreignkey:UserID; not null" json:"user_id"`
	HotelID     uint   `gorm:"foreignkey:HotelID; not null" json:"hotel_id"`
}
