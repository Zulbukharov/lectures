package healthchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckWebsite(t *testing.T) {
	table := []struct {
		name     string
		url      string
		expected bool
	}{
		{"testing valid url", "https://google.com", true},
		{"testing invalid url", "https://invlidurlhaha.com", false},
		{"testing url with 404 status code", "https://google.com/gachi", false},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CheckWebsite(test.url))
		})
	}
}
