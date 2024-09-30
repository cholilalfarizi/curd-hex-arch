package server

import (
	"context"
	"crud-hex/internals/core/services"
	"crud-hex/internals/handlers"
	repo "crud-hex/internals/repositories"
	"crud-hex/internals/utils"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *fiber.App{
	cfg := utils.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Ping to MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//Obtain reference to MongoDB database
	mongoDB := client.Database(cfg.DBName)

	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/db_storage")

	if err != nil{
		log.Fatal(err)
	}
	

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	profilingRepo := repo.NewProfilingRepository(mongoDB)
	profilingService := services.NewProfilingService(profilingRepo)
	productRepo := repo.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := handlers.NewProductController(productService, profilingService)

	app := fiber.New()
	app.Get("/products", productController.FindAll)
	app.Post("/products", productController.Create)
	app.Get("/products/:id", productController.FindByID)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	return app
}