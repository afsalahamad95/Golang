package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type book struct {
	ID     int `gorm:"primaryKey"` // set as primary key in postgres schema
	Title  string
	Author string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "50022021"
	dbname   = "gorm"
)

var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// db parameters

func main() {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// make changes to schema, like you've defined in struct
	db.AutoMigrate(&book{})
	db.Create(&book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// read from db
	var b book
	db.First(&b, 1)
	fmt.Println(b)

	// update the book
	db.Model(&b).Update("Author", "JK Rowling")

	// delete the book
	db.Delete(&b)
}
