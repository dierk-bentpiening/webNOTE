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
	CategoryID string
	CategoryName string
}

type NoteJSON struct {
	ID       	string `json:"id"`
	Title    	string `json:"title"`
	Text 	 	string `json:"text"`
	Author   	string `json:"author"`
	DateTime 	string `json:"DatTime"`
	CategoryID  string `json:"categoryid"`
	CategoryName string `json:"categoryname"`

}
