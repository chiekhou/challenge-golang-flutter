package models

import "time"

type ChatMessage struct {
	ID             uint      `gorm:"primaryKey; not null;" json:"id"`
	GroupeVoyageID uint      `gorm:"not null; default:0" json:"groupe_voyage_id"`
	UserID         uint      `gorm:"not null; default:0" json:"user_id"`
	User           User      `gorm:"foreignKey:UserID" json:"user"`
	Content        string    `gorm:"type:text; not null" json:"content"`
	Created        time.Time `gorm:"default:0" json:"created"`
}
