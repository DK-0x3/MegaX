package handle

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InitDB(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("failed connection db")
		return nil, err
	}
	return db, nil
}
