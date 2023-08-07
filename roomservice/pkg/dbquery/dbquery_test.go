package dbquery

import (
	"errors"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupDb() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return &sqlx.DB{}, errors.New("Failed to load environment variables")
	}
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return &sqlx.DB{}, errors.New("Failed to connect to DB")
	}

	return db, nil
}

func TestDBQuery(t *testing.T) {
	t.Run("should insert new room", func(t *testing.T) {

	})
}
