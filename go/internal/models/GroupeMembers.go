package models

type GroupeMembers struct {
	GroupeVoyageID uint `gorm:"column:groupe_voyage_id; not null"`
	UserID         uint `gorm:"column:user_id; not null"`
}
