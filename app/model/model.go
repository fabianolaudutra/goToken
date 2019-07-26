package model

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
 	"github.com/jinzhu/gorm"
 	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tokens struct {
	gorm.Model
	Token     string     `json:"token"`
	Hash      string     `gorm:"unique" json:"hash"`
	Created_at  *time.Time `json:"created_at"`
	
}
