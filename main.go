package main

import (
	"database/sql"
	"fmt"
	"os"
  "log"
  "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
const(
	user="postgres"
)
func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}	
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	// user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	fmt.Println(psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")
	
	db.Close()
}
