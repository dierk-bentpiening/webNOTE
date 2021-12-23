package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	InternalLibs "WebNote/libs"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		InternalLibs.LogError("Fatal Error: Opening Database failed with Error: " + err.Error())
		panic(err)
	}
}
