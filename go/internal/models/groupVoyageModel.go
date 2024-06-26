package models

type GroupeVoyage struct {
	//gorm.Model
	ID      uint    `gorm:"primary_key; not null" json:"id"`
	Budget  float32 `gorm:"not null" json:"budget"`
	UserID  uint    `gorm:"not null" json:"user_id"`
	Members []User  `gorm:"many2many:group_voyage;" json:"members"`
}
