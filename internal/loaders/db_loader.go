package loaders

import (
	"database/sql"
	"os"

	"github.com/Mohamed-Abdelrazeq/inventory-management-system/internal/database"
	_ "github.com/lib/pq" // add this
)

type DatabaseInstance struct {
	DB *database.Queries
}

func LoadDB() *DatabaseInstance {
	connStr := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("failed to connect database")
	}

	queries := database.New(db)

	return &DatabaseInstance{queries}
}
