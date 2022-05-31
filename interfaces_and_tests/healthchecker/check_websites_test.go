package healthchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://invlidurlhaha.com"
}

func TestCheckWebsitesHealth(t *testing.T) {
	table := []struct {
		name     string
		urls     []string
		expected map[string]bool
	}{
		{"testing urls", []string{"https://google.com", "https://invlidurlhaha.com"}, map[string]bool{
			"https://google.com":        true,
			"https://invlidurlhaha.com": false,
		}},
	}

	for _, test := range table {

		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CheckWebsitesHealth(mockWebsiteChecker, test.urls))
		})
	}
}
