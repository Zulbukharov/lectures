package migrations

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Service struct {
	Migrate *migrate.Migrate
}

func (this *Service) Up() (error, bool) {
	err := this.Migrate.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil, true
		}
		return err, false
	}
	return nil, true
}

func (this *Service) Down() (error, bool) {
	err := this.Migrate.Down()
	if err != nil {
		return err, false
	}
	return nil, true
}

func New(dbConn *sql.DB, migrationsFolderLocation string) (*Service, error) {
	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationsFolderLocation)
	pathToMigrate := strings.Join(dataPath, "")

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, "postgres", driver)
	if err != nil {
		return nil, err
	}
	return &Service{Migrate: m}, nil
}
