package data

import (
	"database/sql"
	"time"
)

type Story struct {
	ID        int
	Title     string
	Content   string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoryModel struct {
	DB *sql.DB
}

func (m StoryModel) Insert(story *Story) error {
	query := `INSERT INTO stories (title, content, username, email) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	return m.DB.QueryRow(query, story.Title, story.Content, story.Username, story.Email).Scan(&story.ID, &story.CreatedAt, &story.UpdatedAt)
}

func (m StoryModel) Get(id int) (*Story, error) {
	story := &Story{}
	query := `SELECT id, title, content, username, email, created_at, updated_at FROM stories WHERE id = $1`
	err := m.DB.QueryRow(query, id).Scan(&story.ID, &story.Title, &story.Content, &story.Username, &story.Email, &story.CreatedAt, &story.UpdatedAt)
	return story, err
}

func (m StoryModel) GetAll() ([]*Story, error) {
	rows, err := m.DB.Query(`SELECT id, title, content, username, email, created_at, updated_at FROM stories ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []*Story
	for rows.Next() {
		s := &Story{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Username, &s.Email, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}
	return stories, nil
}

func (m StoryModel) Update(story *Story) error {
	query := `UPDATE stories SET title = $1, content = $2, username = $3, email = $4, updated_at = now() WHERE id = $5`
	_, err := m.DB.Exec(query, story.Title, story.Content, story.Username, story.Email, story.ID)
	return err
}

func (m StoryModel) Delete(id int) error {
	_, err := m.DB.Exec(`DELETE FROM stories WHERE id = $1`, id)
	return err
}
