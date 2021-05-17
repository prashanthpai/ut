package db

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func NewClient(connStr string) (*DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

type User struct {
	Name  string
	Email string
}

func (db *DB) GetUserByID(ctx context.Context, id int) (*User, error) {

	var user User
	err := db.db.QueryRowContext(ctx, "SELECT name, email FROM users WHERE id = $1", id).Scan(&user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
