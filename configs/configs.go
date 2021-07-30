package configs

import (
	"alterra_store/models/carts"
	"alterra_store/models/productCategories"
	"alterra_store/models/products"
	"alterra_store/models/transactions"
	"alterra_store/models/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configuration {
	var configDB = Configuration{
		DB_USERNAME: "root",
		DB_PASSWORD: "root",
		DB_PORT:     "3306",
		DB_HOST:     "127.0.0.1",
		DB_NAME:     "acp10",
	}
	return configDB
}

func InitDB() {
	configDB := GetConfig()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_USERNAME,
		configDB.DB_PASSWORD,
		configDB.DB_HOST,
		configDB.DB_PORT,
		configDB.DB_NAME)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&productCategories.ProductCategory{})
	DB.AutoMigrate(&products.Product{})
	DB.AutoMigrate(&carts.Cart{})
	DB.AutoMigrate(&transactions.Transaction{})
	DB.AutoMigrate(&transactions.TransactionDetail{})
}

/** TEST */

func GetConfigTest() Configuration {
	var configDB = Configuration{
		DB_USERNAME: "root",
		DB_PASSWORD: "root",
		DB_PORT:     "3306",
		DB_HOST:     "127.0.0.1",
		DB_NAME:     "acp10_test",
	}
	return configDB
}

func InitDBTest() {
	configDB := GetConfigTest()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_USERNAME,
		configDB.DB_PASSWORD,
		configDB.DB_HOST,
		configDB.DB_PORT,
		configDB.DB_NAME)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	MigrationTest()
}

func MigrationTest() {

	DB.Migrator().DropTable(&users.User{})
	DB.AutoMigrate(&users.User{})
}
