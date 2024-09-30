package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
	MongoURI       string
	DBName         string
   }

   func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
	 log.Fatal("Error loading .env file")
	}
   
	return &Config{
	 MongoURI:       os.Getenv("MONGO_URI"),
	 DBName:         os.Getenv("DB_NAME"),
	}
   }