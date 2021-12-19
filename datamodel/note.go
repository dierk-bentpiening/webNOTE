package datamodel

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID string
	Title string
	Text string
	Author string
	DateTime string
}
type NoteJSON struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Text 	 string `json:"text"`
	Author   string `json:"author"`
	DateTime string `json:"DateTime"`
}

type CreatedSuccessFullyJSON struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Title 	 string `json:"title"`
	DateTime string `json:"DateTime"`
}