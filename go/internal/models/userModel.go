package models

type User struct {
	//gorm.Model
	ID           uint           `gorm:"primary_key; not null" json:"id"`
	FirstName    string         `gorm:"size:64; not null; unique" json:"first_name"`
	LastName     string         `gorm:"size:64; not null" json:"last_name"`
	Photo        string         `gorm:"size:255;" json:"photo"`
	Username     string         `gorm:"size:64; not null; unique " json:"username"`
	Password     string         `gorm:"size:255; not null" json:"password"`
	Email        string         `gorm:"size:100; not null; unique" json:"email"`
	Address      string         `gorm:"size:255; not null" json:"address"`
	RoleID       uint           `gorm:"size: 64; default:2; not null" json:"role_id"`
	GroupeVoyage []GroupeVoyage `gorm:"many2many:groupe_members;"json:"groupe_voyage"`
	Voyage       []Voyage       `gorm:"foreignKey:UserId;" json:"voyage"`
}
