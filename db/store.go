package db

import (
	"beta_service/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	models.UserStore
	models.AssetStore
}

func Open() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 "dbuser",
		Passwd:               "dbuserdbuser",
		Net:                  "tcp",
		Addr:                 "baxus.c3tf20wv9p1c.us-east-2.rds.amazonaws.com:3306",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		UserStore:  NewUserStore(db),
		AssetStore: NewAssetStore(db),
	}
}
