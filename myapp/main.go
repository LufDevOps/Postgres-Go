package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
type metrics struct {
	duration *prometheus.SummaryVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		duration: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "tester",
			Name:       "duration_seconds",
			Help:       "Duration of the request.",
			Objectives: map[float64]float64{0.9: 0.01, 0.99: 0.001},
		}, []string{"path", "status"}),
	}
	reg.MustRegister(m.duration)
	return m
}

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Retrieve PostgreSQL connection details from environment variables
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DB")

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

func getDevices(c *fiber.Ctx) error {
	sleep(1000)
	dvs := []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}

	return c.JSON(dvs)
}

func simulateTraffic(m *metrics) {
	for {
		now := time.Now()
		sleep(1000)
		m.duration.WithLabelValues("/fake", "200").Observe(time.Since(now).Seconds())
	}
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}
