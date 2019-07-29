package model

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tokens struct {
	gorm.Model
	Id         int 		  
	Token      string     `json:"token"`
	Hash       string     `gorm:"unique" json:"hash"`
	created_at time.Time  `json:"created_at"`
	updated_at time.Time  
	deleted_at time.Time  
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	
	db.LogMode(true)
	
    db.AutoMigrate(&Tokens{})
	db.Model(&Tokens{})
	return db
}
