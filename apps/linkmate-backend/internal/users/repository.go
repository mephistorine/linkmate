package users

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int
	Email      string
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}

type UserWithPassword struct {
	User
	Password string
}

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
}

type Repository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database: database}
}

func (r *Repository) FindUserById(id int) (*User, error) {
	user := new(User)
	err := r.database.QueryRow("SELECT id, email, name, create_time, update_time FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Email, &user.Name, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) FindUserByEmail(email string) (*User, error) {
	user := new(User)
	err := r.database.QueryRow("SELECT id, email, name, create_time, update_time FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Email, &user.Name, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) FindUserByEmailWithPassword(email string) (*UserWithPassword, error) {
	user := new(UserWithPassword)
	err := r.database.QueryRow("SELECT id, email, name, create_time, update_time, password FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Email, &user.Name, &user.CreateTime, &user.UpdateTime, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) CreateUser(c CreateUserDto) (*User, error) {
	user := new(User)
	err := r.database.QueryRow("INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, create_time, update_time", c.Name, c.Email, c.Password).
		Scan(&user.Id, &user.Name, &user.Email, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) DeleteUserById(id int) error {
	if _, err := r.database.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}
