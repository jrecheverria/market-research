package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID           int
	Sector       int
	Source       string
	Title        string
	URL          string
	Content      string
	Published_At time.Time
	Fetched_At   time.Time
}

type ArticleModel struct {
	DB *sql.DB
}

// TODO What other db operations do we need here?
// Inserts an individual Article
func (a *ArticleModel) Insert(
	sector int,
	source string,
	title string,
	url string,
	content string,
	publishedAt time.Time,
	fetchedAt time.Time) (int, error) {

	stmt := `INSERT INTO articles (sector, source, title, url, content, published_at, fetched_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := a.DB.Exec(stmt, sector, source, title, url, content, publishedAt, fetchedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns an Article based on its most recent id
func (a *ArticleModel) GetById(id int) (*Article, error) {
	return &Article{}, nil
}
