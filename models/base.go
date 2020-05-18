package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rommel96/api-golang-mariadb-echo-jwt/database"
)

var db *gorm.DB

func init() {
	db = database.DBConnect()
}
