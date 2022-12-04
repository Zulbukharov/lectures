package search

import (
	"errors"
)

var (
	ErrMissingArgs   = errors.New("firstName and lastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Searcher interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type PhoneBook struct {
	People []*Person
}

func (p *PhoneBook) Find(searcher Searcher, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := searcher.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}

func (p *PhoneBook) Hello() {}
