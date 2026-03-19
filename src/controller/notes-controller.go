package controller

import (
	"net/http"
	"notes_app/src/service"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	Service *service.NoteService
}

func NewNoteService(s *service.NoteService) *NoteController{
	return  &NoteController{Service: s}
}

func (c *NoteController) NotesGet(ctx *gin.Context){
	res, err := c.Service.GetNotes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Failed get data",
			"code": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(res.Code, gin.H{
		"status": res.Status,
		"message": res.Message,
		"code": res.Code,
		"data": res.Data,
	})
}