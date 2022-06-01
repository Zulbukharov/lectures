package adding_test

import (
	"log"
	"post/internal/app/adding"
	"post/internal/models"
	"post/internal/repository"
	"post/internal/repository/reposuite"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	reposuite.RepositorySuite
}

func (s *repositorySuite) TestAddPostSuccess() {
	adder := adding.New(repository.New(s.DB))
	err := adder.AddPost(models.PostInsert{AuthorID: 2, Content: "hello"})
	if err != nil {
		s.T().Errorf("not expected error got %v", err)
	}
}

func (s *repositorySuite) TestAddPostFailure() {
	adder := adding.New(repository.New(s.DB))
	err := adder.AddPost(models.PostInsert{AuthorID: 0, Content: "hello"})
	if err == nil {
		s.T().Error("expected err got nil")
	}
}

func (s *repositorySuite) SetupTest() {
	log.Println("Starting a Test. Migrating the Database")
	err, _ := s.Migration.Up()
	require.NoError(s.T(), err)
	log.Println("Database Migrated Successfully")
}

func (s *repositorySuite) TearDownTest() {
	log.Println("Finishing Test. Dropping The Database")
	err, _ := s.Migration.Down()
	require.NoError(s.T(), err)
	log.Println("Database Dropped Successfully")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(repositorySuite))
}
