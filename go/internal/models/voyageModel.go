package models

import "time"

type Voyage struct {
	//gorm.Model
	ID          uint       `gorm:"primary_key; not null" json:"id"`
	Date        time.Time  `gorm:"not null" json:"date,omitempty"`
	Destination string     `gorm:"size:64; not null " json:"destination"`
	Activities  []Activity `gorm:"many2many:voyage_activities;" json:"activities"`
}
