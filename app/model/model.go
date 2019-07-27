package model

import (
	"time"

	"github.com/jinzhu/gorm"

)

type Tokens struct {
	gorm.Model
	Token      string     `json:"token"`
	Hash       string     `gorm:"unique" json:"hash"`
	Created_at *time.Time `json:"created_at"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Tokens{})
	db.Model(&Tokens{})
	return db
}
