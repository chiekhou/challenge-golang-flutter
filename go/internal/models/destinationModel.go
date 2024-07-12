package models

type Destination struct {
	//gorm.Model
	ID         uint       `gorm:"primary_key; not null" json:"id"`
	Name       string     `gorm:"size: 255;not null" json:"name"`
	Image      string     `gorm:"size: 255;null" json:"image"`
	Activities []Activity `gorm:"many2many:destination_activities; null" json:"activities"`
	Hotels     []Hotel    `gorm:"many2many:destination_hotels; null" json:"hotels"`
}
