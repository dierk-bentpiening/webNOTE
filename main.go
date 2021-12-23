package main

import "github.com/gin-gonic/gin"
import (
	BaseAPI "WebNote/baseapifunc"
	CategoryFunc "WebNote/categoryfunc"
	DataModel "WebNote/datamodel"
	DatabaseHandler "WebNote/db"
	InternalLibs "WebNote/libs"
	NoteFunc "WebNote/notefunc"
	"log"
	"net/http"
)

var id string

func getIndex(c *gin.Context) {
	InternalLibs.LogInfo("/index called")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "WebNote", "version": "0.0.1-A", "year": "2021"})
}
func main() {
	InternalLibs.LogInfo("Starting Application")
	var err = DatabaseHandler.DB.AutoMigrate(
		&DataModel.Note{},
		&DataModel.Category{},
	)
	if err != nil {
		InternalLibs.LogError(err.Error())
		log.Fatal(err)
	}
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
	router.GET("/category", CategoryFunc.GetCategorys)
	router.POST("/category", CategoryFunc.PostCategory)
	router.Run("localhost:3535")
}
