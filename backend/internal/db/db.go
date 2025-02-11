// backend/internal/db/db.go
package db

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
    connStr := "postgres://user:password@localhost:5432/cmsdb"
    var err error
    DB, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
    }
    fmt.Println("Connected to database!")
}
