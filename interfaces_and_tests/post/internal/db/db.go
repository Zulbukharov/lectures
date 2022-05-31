package db

import (
	"database/sql"
	"fmt"
	"post/internal/config"

	_ "github.com/lib/pq"
)

func New(conf config.Settings) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
