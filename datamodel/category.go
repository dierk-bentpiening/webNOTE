package datamodel

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID string
	Name string
	Tag string
	Author string
	DateTime string
}
type CategoryJSON struct {
	ID       string `json:"id"`
	Name    string `json:"name"`
	Tag 	 string `json:"tag"`
	Author   string `json:"author"`
	DateTime string `json:"DateTime"`
}