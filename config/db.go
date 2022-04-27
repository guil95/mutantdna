package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func GetDBConnection()  *sqlx.DB {
	host := os.Getenv("MUTANT_DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("MUTANT_DB_PORT"))
	user := os.Getenv("MUTANT_DB_USER")
	password := os.Getenv("MUTANT_DB_PASSWORD")
	dbName := os.Getenv("MUTANT_DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping database: %v", err)
		return nil
	}

	return db
}
