package models

import (
	"database/sql"
	"time"
)

// Snippet Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
