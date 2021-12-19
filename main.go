package main

import "github.com/gin-gonic/gin"
import (
	"time"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/google/uuid"
	InternalLibs "WebNote/libs"
	BaseAPI "WebNote/baseapifunc"
	DataModel "WebNote/datamodel"
)
var db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})

var id string

func getIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "WebNote", "version": "0.0.1-A", "year": "2021"})
}

func getNote(c *gin.Context){
	var note DataModel.Note
	notePubID := c.Param("id")
	db.First(&note, "id = ?", notePubID)
	var noteJSON = DataModel.NoteJSON{
		ID: note.ID,
		Title: note.Title,
		Text: note.Text,
		Author: note.Author,
		DateTime: note.DateTime,
	}
	c.IndentedJSON(http.StatusOK, noteJSON)
}

func getNotes(c *gin.Context) {
	var notes []DataModel.Note
	db.Find(&notes)
	var notesJSON []DataModel.NoteJSON
	for _,note := range notes {
		var buffernote = DataModel.NoteJSON{
			ID: note.ID, 
			Title: note.Title, 
			Text: note.Text, 
			Author: note.Author, 
			DateTime: note.DateTime,
		}
		
		notesJSON = append(notesJSON, buffernote)
	}
	c.IndentedJSON(http.StatusOK, notesJSON)
}
func postNote(c *gin.Context) {
	var note DataModel.NoteJSON
	id = uuid.New().String()
	c.BindJSON(&note)
	db.Create(&DataModel.Note{
		ID: id, 
		Title: note.Title, 
		Text: note.Text, 
		Author: note.Author, 
		DateTime: time.Now().String(),
	})
	
	c.IndentedJSON(http.StatusOK, DataModel.CreatedSuccessFullyJSON{
		ID: id,
		Message: "Note Created successfully!",
		Title: note.Title,
		DateTime: time.Now().String(),
	})
}

func main() {	
  	if err != nil {
    	panic("failed to connect database")
  	}
  	db.AutoMigrate(&DataModel.Note{})
	InternalLibs.WelcomeMessage()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", getIndex)
	router.GET("/time", BaseAPI.GetTime)
	router.POST("/note", postNote)
	router.GET("/note", getNotes)
	router.GET("/note/:id", getNote)
	router.Run("localhost:3535")
}