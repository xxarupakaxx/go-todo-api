package domain

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
	News []News `gorm:"many2many:news_topics;"`
}
