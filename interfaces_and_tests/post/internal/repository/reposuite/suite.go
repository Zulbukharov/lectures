package reposuite

import (
	"database/sql"
	"post/internal/config"
	"post/internal/db"
	"post/migrations"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	DB        *sql.DB
	Migration *migrations.Service
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *RepositorySuite) SetupSuite() {
	config, err := config.ParseYAML("../../../config.yml")
	if err != nil {
		suite.T().Fatalf("failed to parse yml: %v", err)
	}
	suite.DB, err = db.New(config)
	if err != nil {
		suite.T().Fatalf("failed to create db connection: %v", err)
	}

	suite.Migration, err = migrations.New(suite.DB, "../../../migrations")
	if err != nil {
		suite.T().Fatalf("failed to create migration: %v", err)
	}

}

func (r *RepositorySuite) TearDownSuite() {
	if r.DB != nil {
		r.DB.Close()
	}
	// if r.m != nil && r.db != nil {
	// 	err := r.m.Down()
	// 	if err != nil {
	// 		log.Fatal("migration failed", err)
	// 	}
	// 	r.db.Close()
	// }
}
