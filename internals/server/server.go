package server

import (
	"context"
	"crud-hex/internals/core/services"
	"crud-hex/internals/handlers"
	repo "crud-hex/internals/repositories"
	"crud-hex/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *fiber.App{
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoDB := client.Database(cfg.DBName)

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDB,
	)

	db, err := sql.Open("mysql", mysqlDSN)

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