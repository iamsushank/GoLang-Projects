package models

import (
	"github.com/iamsushank/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetTopTenBook() []Book {
	var getTenBooks []Book
	db.Model(&Book{}).Limit(10).Order("Name asc").Find(&getTenBooks)
	return getTenBooks
}

func GetBookUsingStruct(books []Book) []Book {
	var getBook []Book
	db.Where(&Book{Name: "Gyz", Author: "NameAS"}).Find(&getBook)
	return getBook
}

func GetBookByAuthorAndName(name string, pub string) []Book {
	var getBookByAuthorAndName []Book
	db.Where("name LIKE ? OR publication=?", "%"+name+"%", pub).Order("Name ASC").Find(&getBookByAuthorAndName)
	return getBookByAuthorAndName

}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id=?", Id).Delete(book)
	return book
}
