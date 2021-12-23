package notefunc

import "github.com/gin-gonic/gin"
import (
	DataModel "WebNote/datamodel"
	DatabaseHandler "WebNote/db"
	Libs "WebNote/libs"
	Responses "WebNote/responses"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var id string

func GetNote(c *gin.Context) {
	var note DataModel.Note
	notePubID := c.Param("id")

	var category DataModel.Category
	DatabaseHandler.DB.First(&note, "id = ?", note.CategoryID)
	DatabaseHandler.DB.First(&note, "id = ?", notePubID)
	var noteJSON = DataModel.NoteJSON{
		ID:         note.ID,
		Title:      note.Title,
		Text:       note.Text,
		Author:     note.Author,
		DateTime:   note.DateTime,
		CategoryID: category.ID,
		CategoryName: category.Name,
	}
	c.IndentedJSON(http.StatusOK, noteJSON)
	Libs.LogInfo("GetNote Called")
}

func GetNotes(c *gin.Context) {
	var notes []DataModel.Note
	DatabaseHandler.DB.Find(&notes)
	var notesJSON []DataModel.NoteJSON
	for _, note := range notes {
		var categoryID string
		var categoryName string
		if note.CategoryID == "NONE" {
			categoryID = "N/A"
			categoryName = "N/A"
		} else {
			var category DataModel.Category
			DatabaseHandler.DB.First(&category, "id = ?", note.CategoryID)
			categoryID = category.ID
			categoryName = category.Name
		}
		var buffernote = DataModel.NoteJSON{
			ID:         note.ID,
			Title:      note.Title,
			Text:       note.Text,
			Author:     note.Author,
			DateTime:   note.DateTime,
			CategoryID: categoryID,
			CategoryName: categoryName,
		}

		notesJSON = append(notesJSON, buffernote)
	}
	c.IndentedJSON(http.StatusOK, notesJSON)
}
func PostNote(c *gin.Context) {
	var note DataModel.NoteJSON
	id = uuid.New().String()
	c.BindJSON(&note)
	var categoryID string
	var categoryName string
	if len(note.CategoryID) == 0 {
		categoryID = "NONE"
		categoryName = "NONE"
	} else {
		var category DataModel.Category
		DatabaseHandler.DB.First(&category, "id = ?", note.CategoryID)
		categoryID = category.ID
		categoryName = category.Name
	}
	DatabaseHandler.DB.Create(&DataModel.Note{
		ID:         id,
		Title:      note.Title,
		Text:       note.Text,
		Author:     note.Author,
		DateTime:   time.Now().String(),
		CategoryID: categoryID,
		CategoryName: categoryName,
	})

	c.IndentedJSON(http.StatusOK, Responses.EntityCreatedSuccessFullyJSON{
		ID:       id,
		Message:  "Note Created successfully!",
		DateTime: time.Now().String(),
	})
}
