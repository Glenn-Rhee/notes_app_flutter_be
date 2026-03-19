package service

import (
	"notes_app/src/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteService struct {
	DB *gorm.DB
}

func NewNoteService(db *gorm.DB) *NoteService{
	return &NoteService{DB: db}
}

func (s *NoteService) GetNotes() (res []model.Note, err error)  {
	var notes []model.Note

	err = s.DB.Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *NoteService) CreateNotes(dataNote model.CreateNote) (res model.NoteResponse, err error){
	id := uint(uuid.New().ID())

	note := model.Note{
		ID: id,
		Title: dataNote.Title,
		Content: dataNote.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := s.DB.Create(&note)

	if result.Error != nil {
		return model.NoteResponse{}, result.Error
	}

	return model.NoteResponse{
		ID: note.ID,
		Title: note.Title,
		Content: note.Content,
		CreatedAt: note.CreatedAt,
	}, nil
}