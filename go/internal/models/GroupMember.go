package models

type GroupMember struct {
	GroupID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}
