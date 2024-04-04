package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

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
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define a route to handle GET requests to /version
	router.GET("/version", getVersion)

	// Define a route to handle liveness probes at /healthz
	router.GET("/healthz", livenessProbe)

	// Define a route to handle readiness probes at /ready
	router.GET("/ready", readinessProbe)

	// Run the Gin router on port 8080
	router.Run(":8080")
}

func getVersion(c *gin.Context) {
	// Query the database for the PostgreSQL version
	var version string
	err := db.QueryRow("SELECT version();").Scan(&version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"version": version})
}

func livenessProbe(c *gin.Context) {
	// Perform a simple check to verify the health of the application
	c.String(http.StatusOK, "OK")
}

func readinessProbe(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
