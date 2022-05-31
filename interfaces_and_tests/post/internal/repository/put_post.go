package repository

import (
	"log"
	"post/internal/models"
)

// Post defines the properties of a Post to be listed
type Post struct {
	ID       uint
	Content  string
	AuthorID uint
}

func (s *repository) Put(u *models.Post) error {
	stmt, err := s.db.Prepare("INSERT INTO posts (content, author_id) VALUES($1, $2);")
	if err != nil {
		log.Printf("db.Prepare: %v", err)
		return ErrFailedToPrepare
	}
	_, err = stmt.Exec(u.Content, u.AuthorID)
	if err != nil {
		return ErrFailedToAdd
	}
	defer stmt.Close()

	return nil
}
