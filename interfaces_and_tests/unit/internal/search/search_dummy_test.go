package search_test

import (
	"testing"
	"unit/internal/search"
)

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*search.Person, firstName, lastName string) *search.Person {
	return &search.Person{}
}

func TestFindReturnsError(t *testing.T) {
	phoneBook := search.PhoneBook{}
	want := search.ErrMissingArgs

	_, got := phoneBook.Find(DummySearcher{}, "", "")

	if got != want {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}
