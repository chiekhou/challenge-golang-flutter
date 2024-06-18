package models

type GroupeVoyage struct {
	ID      uint    `gorm:"primary_key; not null" json:"id"`
	Nom     string  `gorm:"not null" json:"nom"`
	Budget  float32 `gorm:"not null" json:"budget"`
	UserID  uint    `gorm:"not null" json:"user_id"`
	Members []User  `gorm:"many2many:group_voyage;" json:"members"`
}
