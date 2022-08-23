package main

import (
	"fmt"
	"github.com/jmorion/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disabled",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),,
		os.Getenv("APP_DB_NAME")
	)

	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a := App{}
	a.Initialize(db)
	a.Run(":8080")
}