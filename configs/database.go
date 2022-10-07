package configs

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// SetupDB : initializing mysql database
func InitDB() {
	USER := os.Getenv("DBUSERNAME")
	PASS := os.Getenv("DBPASS")
	HOST := os.Getenv("DBHOST")
	PORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	var err error
	DB, err = gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
}
