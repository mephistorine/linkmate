package tags

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"time"
)

type Tag struct {
	Id         int
	Name       string
	Color      string
	UserId     int
	CreateTime time.Time
	UpdateTime time.Time
}

type Repository struct {
	database *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{database: db}
}

type CreateTagDto struct {
	Name   string
	Color  string
	UserId int
}

type UpdateTagDto struct {
	Name  string
	Color string
}

func (r *Repository) Create(dto CreateTagDto) (*Tag, error) {
	tag := new(Tag)

	if err := r.database.QueryRow("INSERT INTO tags(name, color, user_id) VALUES ($1, $2, $3) RETURNING id, name, color, user_id, create_time, update_time", dto.Name, dto.Color, dto.UserId).
		Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime); err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *Repository) UpdateById(id int, dto UpdateTagDto) (*Tag, error) {
	tag := new(Tag)

	if err := r.database.QueryRow("UPDATE tags SET name = $1, color = $2, update_time = NOW() WHERE id = $3 RETURNING id, name, color, user_id, create_time, update_time", dto.Name, dto.Color, id).
		Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime); err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *Repository) DeleteById(id int) error {
	_, err := r.database.Exec("DELETE FROM tags WHERE id = $1", id)
	return err
}

func (r *Repository) FindById(id int) (*Tag, error) {
	tag := new(Tag)

	if err := r.database.QueryRow("SELECT id, name, color, user_id, create_time, update_time FROM tags WHERE id = $1", id).
		Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime); err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *Repository) FindManyByUserId(id int) ([]*Tag, error) {
	rows, err := r.database.Query("SELECT id, name, color, user_id, create_time, update_time FROM tags WHERE user_id = $1", id)

	if err != nil {
		return nil, err
	}

	var tags []*Tag

	defer rows.Close()

	for rows.Next() {
		tag := new(Tag)
		rows.Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime)
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *Repository) FindManyByLinkId(linkId int) ([]*Tag, error) {
	rows, err := r.database.Query("SELECT id, name, color, user_id, create_time, update_time FROM tags RIGHT JOIN public.links_tags lt ON tags.id = lt.tag_id WHERE lt.link_id = $1", linkId)

	if err != nil {
		return nil, err
	}

	var tags []*Tag

	defer rows.Close()

	for rows.Next() {
		tag := new(Tag)
		rows.Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime)
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *Repository) FindManyByLinkIds(linkIds []int) (map[int][]*Tag, error) {
	rows, err := r.database.Query("SELECT id, name, color, user_id, create_time, update_time, lt.link_id FROM tags RIGHT JOIN public.links_tags lt ON tags.id = lt.tag_id WHERE lt.link_id = ANY ($1)", pq.Array(linkIds))

	if err != nil {
		return nil, err
	}

	var tagsByLinkId = make(map[int][]*Tag)

	defer rows.Close()

	for rows.Next() {
		var linkId int
		tag := new(Tag)
		rows.Scan(&tag.Id, &tag.Name, &tag.Color, &tag.UserId, &tag.CreateTime, &tag.UpdateTime, &linkId)
		tagsByLinkId[linkId] = append(tagsByLinkId[linkId], tag)
	}

	return tagsByLinkId, nil
}

func (r *Repository) CreateTagSettings(linkId int, tagIds []int) error {
	var values []string

	for _, id := range tagIds {
		values = append(values, fmt.Sprintf("(%d, %d)", linkId, id))
	}

	if _, err := r.database.Exec(fmt.Sprintf("INSERT INTO links_tags (link_id, tag_id) VALUES %s", strings.Join(values, ","))); err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTagSettings(linkId int, tagIds []int) error {
	_, err := r.database.Exec("DELETE FROM links_tags WHERE link_id = $1 AND tag_id = ANY ($2)", linkId, pq.Array(tagIds))
	return err
}
