package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {

	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "parseTime=true&loc=Asia%2fJakarta")

	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	DB = db
}