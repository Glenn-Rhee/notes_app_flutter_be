package service

// Import library needs
import (
	"net/http"
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

func (s *NoteService) DeleteNotes(id uint) (string, int) {
	err := s.DB.First(&model.Note{}, id).Error

	if err != nil {
		return "Data note not found!", http.StatusNotFound
	}

	err = s.DB.Delete(&model.Note{}, id).Error
	if err != nil {
		return "Failed delete data note!", http.StatusInternalServerError
	}

	return "", http.StatusOK
}

func (s *NoteService) UpdateNotes(id uint, dataNote model.CreateNote) (string, int) {
	result := s.DB.First(&model.Note{}, id)

	if result.Error != nil {
		return "Data note not found!", http.StatusNotFound
	}

	result = s.DB.Model(&model.Note{}).
				Where("id = ?", id).
				Updates(model.Note{
					Title:   dataNote.Title,
					Content: dataNote.Content,
					UpdatedAt: time.Now(),
				})
	
	if result.Error != nil {
		return "Failed update data note!", http.StatusInternalServerError
	}

	return "", http.StatusOK
}