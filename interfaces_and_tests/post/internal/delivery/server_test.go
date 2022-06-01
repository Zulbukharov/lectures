package delivery_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"post/internal/app/adding"
	"post/internal/delivery"
	"post/internal/models"
	"testing"
)

type StubAddingStore struct {
	ret error
}

func (s *StubAddingStore) AddPost(u models.PostInsert) error {
	return s.ret
}

func TestAddPost(t *testing.T) {

	tests := []struct {
		desc   string
		server *delivery.Server
		code   int
		body   []byte
	}{
		{
			desc:   "it returns 200 on /post with valid request",
			server: delivery.NewServer(&StubAddingStore{nil}),
			code:   http.StatusOK,
			body:   []byte(`{"author_id": 1, "content": "hello"}`),
		},
		{
			desc:   "it returns 400 on /post with invalid data in request",
			server: delivery.NewServer(&StubAddingStore{adding.ErrInvalidInput}),
			code:   http.StatusBadRequest,
			body:   []byte(`{"author_id":0, "content": "hello"}`),
		},
		{
			desc:   "it returns 400 on /post with broken data in request",
			server: delivery.NewServer(&StubAddingStore{adding.ErrInvalidInput}),
			code:   http.StatusBadRequest,
			body:   []byte(`{"author_id":, "content": "hello"}`),
		},
	}

	for _, test := range tests {
		t.Run("it returns 200 on /post", func(t *testing.T) {
			body := bytes.NewBuffer(test.body)
			defer body.Reset()

			request, _ := http.NewRequest(http.MethodPost, "/post", body)
			response := httptest.NewRecorder()
			test.server.ServeHTTP(response, request)

			var got delivery.Response

			err := json.NewDecoder(response.Body).Decode(&got)
			if err != nil {
				t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
			}

			if response.Code != test.code {
				t.Fatalf("Expected %v status code got %v status code", response.Code, test.code)
			}
		})
	}

}
