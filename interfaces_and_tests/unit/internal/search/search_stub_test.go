package search_test

import (
	"testing"
	"unit/internal/search"
)

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*search.Person, firstName, lastName string) *search.Person {
	return &search.Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+77773080226"
	phoneBook := &search.PhoneBook{}

	phone, _ := phoneBook.Find(StubSearcher{fakePhone}, "Abylaikhan", "Zulbukharov")

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}
}
