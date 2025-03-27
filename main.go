package main

import (
	"mpm_server/mpm_server/models"
	"mpm_server/mpm_server/routes"
	mpmserverfrontend "mpm_server/mpm_server_frontend"
	"os"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loding .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	models.InitDB(config)
	routes.AuthRoutes(r)
	mpmserverfrontend.LoadHTMLs(r)
	mpmserverfrontend.LoadCSSs(r)

	err = r.RunTLS(":8080", "ssl_stuff/cert.pem", "ssl_stuff/private.key")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}
