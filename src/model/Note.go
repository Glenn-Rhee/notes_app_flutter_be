package model

import "time"

type Note struct {
	ID        uint		`gorm:"primaryKey"`
	Title     string	`gorm:"type:varchar(255);not null"`
	Content   string	`gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}