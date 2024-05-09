package models

import "gorm.io/gorm"

type GroupeVoyage struct {
	gorm.Model
	Budget    float32 `gorm:"not null" json:"budget"`
	Personnes string  `gorm:"not null" json:"personnes"`
	Roadmap   string  `gorm:"not null" json:"roadmap"`
}
