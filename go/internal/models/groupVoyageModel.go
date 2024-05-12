package models

import "gorm.io/gorm"

type GroupeVoyage struct {
	gorm.Model
	Budget  float32 `gorm:"not null" json:"budget"`
	Roadmap string  `gorm:"not null" json:"roadmap"`
	UserID  uint    `gorm:"not null" json:"user_id"`
}
