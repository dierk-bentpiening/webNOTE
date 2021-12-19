package main

import "github.com/gin-gonic/gin"
import (
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	InternalLibs "WebNote/libs"
	BaseAPI "WebNote/baseapifunc"
	DataModel "WebNote/datamodel"
	NoteFunc "WebNote/notefunc"

)
var db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})

var id string

func getIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "WebNote", "version": "0.0.1-A", "year": "2021"})
}
func main() {
	InternalLibs.LogInfo("Starting Application")	
  	if err != nil {
    	panic("failed to connect database")
  	}

  	db.AutoMigrate(&DataModel.Note{})
  	db.AutoMigrate(&DataModel.Category{})
	InternalLibs.LogInfo("Migrated DataModel ")
	InternalLibs.WelcomeMessage()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", getIndex)
	router.GET("/time", BaseAPI.GetTime)
	router.POST("/note", NoteFunc.PostNote)
	router.GET("/note", NoteFunc.GetNotes)
	router.GET("/note/:id", NoteFunc.GetNote)
	router.Run("localhost:3535")
}