package notefunc

import "github.com/gin-gonic/gin"
import (
	"time"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/google/uuid"
	DataModel "WebNote/datamodel"
	Libs "WebNote/libs"
	Responses "WebNote/responses"

)
var db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})

var id string

func GetNote(c *gin.Context){
	var note DataModel.Note
	notePubID := c.Param("id")

	var category DataModel.Category
	db.First(&note, "id = ?", note.CategoryID)
	db.First(&note, "id = ?", notePubID)
	var noteJSON = DataModel.NoteJSON{
		ID: note.ID,
		Title: note.Title,
		Text: note.Text,
		Author: note.Author,
		DateTime: note.DateTime,
		CategoryID: category.Name,
	}
	c.IndentedJSON(http.StatusOK, noteJSON)
	Libs.LogInfo("GetNote Called")
}

func GetNotes(c *gin.Context) {
	var notes []DataModel.Note
	db.Find(&notes)
	var notesJSON []DataModel.NoteJSON
	for _,note := range notes {
		
		var category DataModel.Category
		db.First(&note, "id = ?", note.CategoryID)
		var buffernote = DataModel.NoteJSON{
			ID: note.ID, 
			Title: note.Title, 
			Text: note.Text, 
			Author: note.Author, 
			DateTime: note.DateTime,
			CategoryID: category.Name,

		}
		
		notesJSON = append(notesJSON, buffernote)
	}
	c.IndentedJSON(http.StatusOK, notesJSON)
}
func PostNote(c *gin.Context) {
	var note DataModel.NoteJSON
	id = uuid.New().String()
	c.BindJSON(&note)
	db.Create(&DataModel.Note{
		ID: id, 
		Title: note.Title, 
		Text: note.Text, 
		Author: note.Author, 
		DateTime: time.Now().String(),
		CategoryID: note.CategoryID,
	})
	
	c.IndentedJSON(http.StatusOK, Responses.EntityCreatedSuccessFullyJSON{
		ID: id,
		Message: "Note Created successfully!",
		Title: note.Title,
		DateTime: time.Now().String(),
	})
}