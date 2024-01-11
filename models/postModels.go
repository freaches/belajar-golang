package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string
	Body   string
	TestID uint
	Test   Test `json:"test"`
}
