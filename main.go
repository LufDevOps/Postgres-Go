package main

import (
	"database/sql"
	"fmt"
	// "log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }

	// Retrieve PostgreSQL connection details from environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_DBPASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")

	// Construct connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the PostgreSQL database
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define a route to handle GET requests to /version
	router.GET("/version", func(c *gin.Context) {
		// Query the database for the PostgreSQL version
		var version string
		err := conn.QueryRow("SELECT version();").Scan(&version)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"version": version})
	})

	// Run the Gin router on port 8080
	router.Run(":8080")
}
