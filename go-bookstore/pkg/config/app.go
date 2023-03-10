package config

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost user=postgres password=root dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//d, err := gorm.Open("mysql", "root:root/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
