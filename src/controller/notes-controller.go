package controller

import (
	"log"
	"net/http"
	"notes_app/src/model"
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
		log.Printf("Error while get data notes: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"message": "Failed get data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Successfully get data notes!",
		"data": res,
	})
}

func (c *NoteController) NotesPost(ctx *gin.Context){
	var reqBody model.CreateNote

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Failed get data",
			"code": http.StatusBadRequest,
		})
		return
	}

	res, err := c.Service.CreateNotes(reqBody)

	if err != nil {
		log.Printf("Failed create data note: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"message": "Failed create data",
			"code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"message": "Successfully create data note",
		"data": res,
	})
}

func (c *NoteController) NotesDelete(ctx *gin.Context){
	
}