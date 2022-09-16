package main

import (
	"flag"
	"fmt"
	"github.com/jmorion/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	createUsersQueue = "CREATE_USER"
	updateUsersQueue = "UPDATE_USER"
	deleteUsersQueue = "DELETE_USER"
)

func main() {
	cache := Cache(Enable: true)

	flag.StringVar(
		&cache.Address,
		"redis_address",
		os.Getenv("APP_RD_ADDRESS"),
		"Redis Address"
	)

	flag.StringVar(
		&cache.Auth,
		"redis_auth",
		os.Getenv("APP_RD_AUTH"),
		"Redis Auth"
	)

	flag.StringVar(
		&cache.DB,
		"redis_db_name",
		os.Getenv("APP_RD_DBNAME"),
		"Redis DB name"
	)

	flag.IntVar(
		&cache.MaxIdle,
		"redis_max_idle",
		10,
		"Redis max Idle"
	)

	flag.IntVar(
		&cache.MaxActive,
		"redis_max_active",
		100,
		"Redis max Active"
	)

	flag.IntVar(
		&cache.IdleTimeoutSecs,
		"redis_timeout",
		60,
		"Redis timeout in seconds"
	)

	flag.intVar(
		&numWorkers,
		"num_workers",
		10,
		"Number of workers to consume queue"
	)

	flag.Parse()

	cache.Pool = cache.NewCachePool()

	connectionString := os.Getenv("DATABASE_DEV_URL")

	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	go UsersToDB(numWorkers, db. cache, createUsersQueue)
	go UsersToDB(numWorkers, db. cache, updateUsersQueue)
	go UsersToDB(numWorkers, db. cache, deleteUsersQueue)

	a := App{}
	a.Initialize(cache, db)
	a.Run(":3000")
}