package server

import (
	"crud-hex/internals/core/services"
	"crud-hex/internals/handlers"
	mySqlRepo "crud-hex/internals/repositories"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App{
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/db_storage")

	if err != nil{
		log.Fatal(err)
	}
	

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := mySqlRepo.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := handlers.NewProductController(productService)

	app := fiber.New()
	app.Get("/products", productController.FindAll)
	app.Post("/products", productController.Create)
	app.Get("/products/:id", productController.FindByID)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	return app
}