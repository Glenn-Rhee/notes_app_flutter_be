package model

import "time"

// Note struct untuk merepresentasikan data note dalam database
type Note struct {
	ID        uint		`gorm:"primaryKey"`
	Title     string	`gorm:"type:varchar(255);not null"`
	Content   string	`gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Struct untuk request body saat membuat note baru
type CreateNote struct {
	Title 	string `json:"title"`
	Content string `json:"content"`
}

type NoteResponse struct {
	ID 		uint
	Title 	string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time
}