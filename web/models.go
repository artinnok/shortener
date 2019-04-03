package web

import "github.com/jinzhu/gorm"

type URLModel struct {
	gorm.Model
	LongLink  string
	ShortLink string
	Clicks    int
}
