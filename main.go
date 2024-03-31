package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)
var db *gorm.DB
var err error

func main() {
	// Loading enviroment variables
	/*dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")*/
	
	dialect :="postgres"
	host :="localhost"
	dbPort :="5432"
	user :="postgres"
	dbpassword :="giang2002"


	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s sslmode=disable password=%s port=%s", host, user, dbpassword, dbPort)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully ")

	// Close the databse connection when the main function closes
	defer db.Close()

}