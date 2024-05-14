package models

import (
	"time"
)

type Destination struct {
	//gorm.Model
	ID              uint      `gorm:"primary_key; not null" json:"id"`
	Departure       time.Time `gorm:"not null" json:"departure"`
	Return          time.Time `gorm:"not null" json:"return"`
	NameDestination string    `gorm:"size: 255;not null" json:"name_destination"`
	Type            string    `json:"type; not null" json:"type"`
	UserID          uint      `gorm:"foreignkey:UserID; not null" json:"user_id"`
}
