package models

import (
	"time"
)

type Voyage struct {
	//gorm.Model
	ID          uint       `gorm:"primary_key; not null" json:"id"`
	DateAller   time.Time  `gorm:"not null" json:"dateAller,omitempty"`
	DateRetour  time.Time  `gorm:"not null" json:"dateRetour,omitempty"`
	Destination string     `gorm:"size:64; not null " json:"destination"`
	Activities  []Activity `gorm:"many2many:voyage_activities;" json:"activities"`
	Hotels      []Hotel    `gorm:"many2many:voyage_hotels; null" json:"hotels"`
	UserId      uint       `gorm:"not null" json:"user_id"`
}
