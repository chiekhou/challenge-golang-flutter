package models

import (
	"gorm.io/gorm"
	"time"
)

type Destination struct {
	gorm.Model
	Departure       time.Time `gorm:"not null" json:"departure"`
	Return          time.Time `gorm:"not null" json:"return"`
	NameDestination string    `gorm:"size: 255;not null" json:"name_destination"`
	Type            string    `json:"type; not null" json:"type"`
}
