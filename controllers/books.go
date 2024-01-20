package controllers

import (
	"crud-with-postgres/configs"
	"crud-with-postgres/database"
	"crud-with-postgres/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// create
func CreateBook(c *fiber.Ctx) error {
	db := database.DB
	body := new(models.Book)

	err := c.BodyParser(body)
	if err != nil {
		return configs.Error(c, "Error Validating Body", err)
	}

	fmt.Printf("\n\nGot here with body %v", db)
	book := models.Book{Name: body.Name, Author: body.Author}

	data := db.Create(&book)
	log.Println(data)
	// result := APIUser{ID: data}
	return configs.Success(c, "Books created successfully", body)
}

// update
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.Book)

	err := c.BodyParser(body)
	if err != nil {
		return configs.Error(c, "Error Validating Body", err)
	}

	book := models.Book{}
	bookExist := database.DB.First(&book, id).Error

	log.Println(book)
	if bookExist != nil {
		return configs.Error(c, "Book does not exist", bookExist)
	}

	intNumber, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
	}
	database.DB.Save(&models.Book{ID: uint(intNumber), Name: body.Name})
	return configs.Success(c, "Book updated successfully", book)
}

// find one
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book := models.Book{}
	database.DB.First(&book, id)

	return configs.Success(c, "Book retrieved", book)
}

// delete
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book := models.Book{}
	bookExist := database.DB.First(&book, id)

	if bookExist.Error != nil {
		return configs.Error(c, "Book not found", bookExist.Error)
	}

	database.DB.Delete(&book)

	return configs.Success(c, "Record Deleted successfully", book)
}

type APIUser struct {
	ID   uint
	Name string
}

// find all
func FetchBook(c *fiber.Ctx) error {
	limit := c.Query("limit")
	books := []models.Book{}
	intNumber, err := strconv.Atoi(limit)

	if err != nil {
		log.Fatal(err)
	}

	database.DB.Limit(intNumber).Find(&books)

	return configs.Success(c, "Books retrieved", books)
}
