package repository_test

import (
	"errors"
	"post/internal/models"
	"post/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPutPost(t *testing.T) {
	tests := []struct {
		desc        string
		mockClosure func(mock sqlmock.Sqlmock)
		args        models.PostInsert
		expected    error
	}{
		{
			desc: "success put",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO posts").
					ExpectExec().
					WithArgs("hello", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args:     models.PostInsert{AuthorID: 1, Content: "hello"},
			expected: nil,
		},
		{
			desc: "failure on put",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO posts").
					ExpectExec().
					WithArgs("hello", 1).
					WillReturnError(repository.ErrFailedToAdd)
			},
			args:     models.PostInsert{AuthorID: 1, Content: "hello"},
			expected: repository.ErrFailedToAdd,
		},
		{
			desc: "failed put",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO posts").WillReturnError(errors.New("hey"))
			},
			args:     models.PostInsert{AuthorID: 1, Content: "hello"},
			expected: repository.ErrFailedToPrepare,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			handler, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			test.mockClosure(mock)

			db := repository.New(handler)
			if err = db.Put(&test.args); err != test.expected {
				t.Errorf("was expected %v while putting post, got %v", test.expected, err)
			}

			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = handler.Close(); err != nil {
				t.Error(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})
	}
}
