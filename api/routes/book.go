package routes

import (
	"cmc-web-scraper/api/handlers"
	"cmc-web-scraper/pkg/book"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
