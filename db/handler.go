package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	InternalLibs "WebNote/libs"
)

var DB *gorm.DB

func init() {
	var cfg = InternalLibs.GetConfigValues()
	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.Database.DBName), &gorm.Config{})
	if err != nil {
		InternalLibs.LogError("Fatal Error: Opening Database failed with Error: " + err.Error())
		panic(err)
	}
}
