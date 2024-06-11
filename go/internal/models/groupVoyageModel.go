package models

import (
	"time"
)

type GroupeVoyage struct {
	ID            uint      `gorm:"primary_key; not null" json:"id"`
	Budget        float32   `gorm:"not null" json:"budget"`
	Roadmap       string    `gorm:"not null" json:"roadmap"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	NbPersonnes   int       `gorm:"not null" json:"nb_personnes"`
	DateDepart    time.Time `gorm:"not null" json:"date_depart"`
	DateRetour    time.Time `gorm:"not null" json:"date_retour"`
	Nom           string    `gorm:"not null" json:"nom"`
	DestinationID uint      `gorm:"not null" json:"destination_id"`
}
