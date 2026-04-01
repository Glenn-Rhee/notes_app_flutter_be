package main

import (
	"log"
	"notes_app/src/controller"
	"notes_app/src/lib"
	"notes_app/src/model"
	"notes_app/src/service"

	"github.com/gin-gonic/gin"
)

func main(){
	db, err := lib.DbConnect()
	if err != nil {
		log.Fatalf("Failed to connect DB: %v\n", err)
	}

	db.AutoMigrate(&model.Note{})

	noteService := service.NewNoteService(db)
	noteController := controller.NewNoteService(noteService)

	router := gin.Default()

	router.GET("/notes", noteController.NotesGet)
	router.POST("/notes", noteController.NotesPost)
	router.DELETE("/notes", noteController.NotesDelete)
	
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}