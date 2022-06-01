package adding_test

import (
	"errors"
	"post/internal/app/adding"
	"post/internal/models"
	"post/internal/repository"

	"post/internal/repository/mock"
	"testing"

	"github.com/gojuno/minimock/v3"
)

func TestAddPost(t *testing.T) {
	tests := []struct {
		desc        string
		mockClosure func(*mock.RepositoryMock)
		args        models.PostInsert
		expected    error
	}{
		{
			desc: "success add post",
			mockClosure: func(rm *mock.RepositoryMock) {
				rm.PutMock.Return(nil)
			},
			args:     models.PostInsert{AuthorID: 1, Content: "hello"},
			expected: nil,
		},
		{
			desc: "failure add post",
			mockClosure: func(rm *mock.RepositoryMock) {
				rm.PutMock.Return(repository.ErrFailedToAdd)
			},
			args:     models.PostInsert{AuthorID: 1, Content: "hello"},
			expected: adding.ErrFailedToAdd,
		},
		{
			desc: "failure add post on validation",
			mockClosure: func(rm *mock.RepositoryMock) {

			},
			args:     models.PostInsert{AuthorID: 0, Content: "hello"},
			expected: adding.ErrInvalidInput,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Finish()
			mockRepo := mock.NewRepositoryMock(mc)
			test.mockClosure(mockRepo)
			service := adding.New(mockRepo)
			if err := service.AddPost(test.args); !errors.Is(err, test.expected) {
				t.Errorf("was expected %v while adding post, got %v", test.expected, err)
			}
		})
	}

}
