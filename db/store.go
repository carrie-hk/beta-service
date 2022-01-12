package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*AssetStore
	*UserStore
}

func NewStore() (*Store, error) {

	//I need to put into an env file
	cfg := mysql.Config{
		User:                 "dbuser",
		Passwd:               "dbuserdbuser",
		Net:                  "tcp",
		Addr:                 "baxus.c3tf20wv9p1c.us-east-2.rds.amazonaws.com:3306",
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		AssetStore: &AssetStore{DB: db},
		UserStore:  &UserStore{DB: db},
	}, nil
}
