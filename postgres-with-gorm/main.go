package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}
type Book struct {
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBook)
	api.Get("/get_books", r.GetBooks)

}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	context.BodyParser(&book)
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	r := Repository{DB: db}
	r.SetupRoutes(app)
	app.Listen(":4000")
}
