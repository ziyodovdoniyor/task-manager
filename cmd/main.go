package main

import (
	"alfa1/config"
	"alfa1/postgres"
)

func main() {
	db, err := config.Configuration()
	if err != nil {
		panic(err)
	}

	_, err = postgres.Connection(db)
	if err != nil {
		panic(err)
	}

}
