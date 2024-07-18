package models

type GroupeVoyage struct {
	ID      uint          `gorm:"primary_key; not null" json:"id"`
	Nom     string        `gorm:"not null" json:"nom"`
	Budget  float32       `gorm:"not null" json:"budget"`
	UserID  uint          `gorm:"not null" json:"user_id"`
	User    User          `gorm:"foreignkey:UserID" json:"-"`
	Members []User        `gorm:"many2many:groupe_members;OnUpdate:CASCADE,OnDelete:SET NULL;" json:"members"`
	Chats   []ChatMessage `gorm:"foreignKey:GroupeVoyageID;OnUpdate:CASCADE,OnDelete:SET NULL;" json:"chats"`
	Voyage  Voyage        `gorm:"foreignkey:GroupeVoyageID;OnUpdate:CASCADE,OnDelete:SET NULL;" json:"voyage"`
}
