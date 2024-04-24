package models

import "gorm.io/gorm"

type GroupeVoyage struct {
	gorm.Model
	Budget    float32 `json:"budget"`
	Personnes string  `json:"personnes"`
	Roadmap   string  `json:"roadmap"`
}
