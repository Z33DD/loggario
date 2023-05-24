package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	ID   int    `json: "object_id"`
	Name string `json: "name"`
}
