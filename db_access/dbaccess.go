package db_access

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbAccess struct {
	*AssetDbAccess
	*UserDbAccess
}

func NewDbAccess() (*DbAccess, error) {

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
