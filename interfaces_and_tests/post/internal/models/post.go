package models

// Post basic post struct
type Post struct {
	ID       int32  `json:"id"`
	AuthorID uint   `json:"author_id"`
	Content  string `json:"content"`
}

type PostInsert struct {
	AuthorID uint   `json:"author_id"`
	Content  string `json:"content"`
}
