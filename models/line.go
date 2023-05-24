package models

import "gorm.io/gorm"

type Line struct {
	gorm.Model
	CategoryID int `json: "category_id"`
	Category   Category
	ObjectID   int `json: "object_id"`
	Object     Object
	Content    string `json: "content"`
}
