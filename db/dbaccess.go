package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbAccess struct {
	*AssetDbAccess
	*UserDbAccess
}

func NewDbAccess() (*DbAccess, error) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//I need to put into an env file
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_ADDR"),
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &DbAccess{
		AssetDbAccess: &AssetDbAccess{DB: db},
		UserDbAccess:  &UserDbAccess{DB: db},
	}, nil
}
