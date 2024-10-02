package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
	MongoURI       string
	DBName         string
	MySQLUser   string
 	MySQLPassword   string
 	MySQLHost   string
 	MySQLPort   string
 	MySQLDB string
   }

   func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
	 log.Fatal("Error loading .env file")
	}
   
	return &Config{
	 MongoURI:       os.Getenv("MONGO_URI"),
	 DBName:         os.Getenv("DB_NAME"),
	 MySQLUser: os.Getenv("MYSQL_USER"),
	 MySQLPassword: os.Getenv("MYSQL_PASSWORD"),
	 MySQLHost: os.Getenv("MYSQL_HOST"),
	 MySQLPort: os.Getenv("MYSQL_PORT"),
	 MySQLDB: os.Getenv("MYSQL_DB"),
	}
   }