package mysql

import (
	"ginarticle/model"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Dbinit() (db *gorm.DB) {
	db, err = gorm.Open("mysql", "root:.aA1451418@tcp(123.207.88.76:3306)/ginarticle?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	return db

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
