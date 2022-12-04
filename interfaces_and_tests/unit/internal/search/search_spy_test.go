package search_test

import (
	"testing"
	"unit/internal/search"
)

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	chain           []string
}

func (ss *SpySearcher) Search(people []*search.Person, firstName, lastName string) *search.Person {
	ss.searchWasCalled = true
	ss.chain = append(ss.chain, "Search")
	return &search.Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+77773080226"
	phoneBook := &search.PhoneBook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phoneBook.Find(spy, "Abylaikhan", "Zulbukharov")

	if !spy.searchWasCalled {
		t.Errorf("expected to call 'Search' in 'Find', but it wasn't.")
	}

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}
}
