package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/acp10?charset=utf8mb4"

	var error error

	DB, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
