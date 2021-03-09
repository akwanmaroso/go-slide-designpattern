package main

import (
	"log"
)

// SINGLETON OMIT
type database struct {
}

var db *database

func GetConnection() *database {
	if db == nil {
		db = new(database)
		log.Println("Create connection to db...")
	}
	log.Println("Connection already exists...")
	return db
}

// END OMIT

// MAIN OMIT
func main() {
	for i := 0; i < 10; i++ {
		GetConnection()
	}
}

// END OMIT
