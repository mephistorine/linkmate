package links

import (
	"database/sql"
	"time"
)

type CreateLinkDto struct {
	Key    string
	Url    string
	UserId int
}

type Link struct {
	Id         int
	Key        string
	Url        string
	UserId     int
	CreateTime time.Time
}

type Repository struct {
	database *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{database: db}
}

func (r *Repository) Create(dto CreateLinkDto) (*Link, error) {
	link := new(Link)

	if err := r.database.QueryRow("INSERT INTO links(key, url, user_id) VALUES($1, $2, $3) RETURNING id, key, url, user_id, create_time", dto.Key, dto.Url, dto.UserId).
		Scan(&link.Id, &link.Key, &link.Url, &link.UserId, &link.CreateTime); err != nil {
		return nil, err
	}

	return link, nil
}

func (r *Repository) DeleteById(id int) error {
	_, err := r.database.Exec("DELETE FROM links WHERE id = $1", id)
	return err
}

func (r *Repository) FindOneById(id int) (*Link, error) {
	link := new(Link)

	if err := r.database.QueryRow("SELECT id, key, url, user_id, create_time FROM links WHERE id = $1", id).
		Scan(&link.Id, &link.Key, &link.Url, &link.UserId, &link.CreateTime); err != nil {
		return nil, err
	}

	return link, nil
}

func (r *Repository) FindOneByKey(key string) (*Link, error) {
	link := new(Link)

	if err := r.database.QueryRow("SELECT id, key, url, user_id, create_time FROM links WHERE key = $1", key).
		Scan(&link.Id, &link.Key, &link.Url, &link.UserId, &link.CreateTime); err != nil {
		return nil, err
	}

	return link, nil
}

func (r Repository) FindManyByUserId(userId int) ([]*Link, error) {
	var links []*Link
	rows, err := r.database.Query("SELECT id, key, url, user_id, create_time FROM links WHERE user_id = $1", userId)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		link := new(Link)
		rows.Scan(&link.Id, &link.Key, &link.Url, &link.UserId, &link.CreateTime)
		links = append(links, link)
	}

	return links, nil
}
