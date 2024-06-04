package models

type User struct {
	//gorm.Model
	ID        uint   `gorm:"primary_key; not null" json:"id"`
	FirstName string `gorm:"size:64; not null; unique" json:"first_name"`
	LastName  string `gorm:"size:64; not null" json:"last_name"`
	Username  string `gorm:"size:64; not null; unique " json:"username"`
	Password  string `gorm:"size:255; not null" json:"password"`
	Email     string `gorm:"size:100; not null; unique" json:"email"`
	Address   string `gorm:"size:255; not null" json:"address"`
	RoleID    uint   `gorm:"size: 64; default:1; not null" json:"role_id"`
	//Role        Role         `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	GroupVoyage GroupeVoyage `json:"group_voyage"`
}
