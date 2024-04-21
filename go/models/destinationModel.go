package models

import (
	"gorm.io/gorm"
	"time"
)

type Destination struct {
	gorm.Model
	Departure       time.Time `gorm:"type:varchar(30);not null"`
	Return          time.Time `gorm:"type:varchar(30);not null"`
	NameDestination string    `json:"name_destination"`
	Type            string    `json:"type"`
}
