package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *  gorm.DB
)

func Connect() {
    dsn := "root:andoriyaprashant@tcp(localhost:3306)/ANDORIYAPRASHANT?charset=utf8&parseTime=True&loc=Local"
    d, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB{
	return db
}