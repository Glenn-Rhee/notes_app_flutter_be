package service

import (
	"net/http"
	"notes_app/src/model"

	"gorm.io/gorm"
)

type NoteService struct {
	DB *gorm.DB
}

func NewNoteService(db *gorm.DB) *NoteService{
	return &NoteService{DB: db}
}

func (s *NoteService) GetNotes() (res model.ResponsePayload, err error)  {
	var notes []model.Note

	err = s.DB.Find(&notes).Error
	return model.ResponsePayload{
		Code: http.StatusOK,
		Status: "Success",
		Message: "Successfully get data",
		Data: notes,
	}, nil
}