package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("valid test", func(t *testing.T) {
		slowServer := makeServer(time.Millisecond * 20)
		fastServer := makeServer(time.Millisecond * 0)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("test with timeout", func(t *testing.T) {
		slowServer := makeServer(time.Millisecond * 500)
		fastServer := makeServer(time.Millisecond * 600)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, time.Millisecond*10)
		if err == nil {
			t.Errorf("got nil, want err")
		}
	})

}

func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
