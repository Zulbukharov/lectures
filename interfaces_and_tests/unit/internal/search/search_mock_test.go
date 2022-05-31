package search_test

import (
	"testing"
	"unit/internal/search"
)

type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*search.Person, firstName, lastName string) *search.Person {
	ms.methodsToCall["Search"] = true
	return &search.Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}
	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+77773080226"
	phoneBook := &search.PhoneBook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phoneBook.Find(mock, "Abylaikhan", "Zulbukharov")

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}
