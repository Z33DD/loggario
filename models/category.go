package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   int    `json: "category_id"`
	Name string `json: "name"`
}
