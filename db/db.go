package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

// データベースをオープンします。
func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// データベースをクローズします。
func (d *Database) Close() error {
	return d.db.Close()
}

// テーブルを生成します。
func (d *Database) CreateTables() error {
	// users
	_, err := d.db.Exec(`create table users (
id integer primary key autoincrement not null,
user_name text not null,
password_hash NONE,
unique (user_name)
);`)
	if err != nil {
		return err
	}

	// posts
	_, err = d.db.Exec(`create table posts (
		id integer primary key autoincrement not null,
		user_id integer not null,
		message text,
		url text
		);`)
	if err != nil {
		return err
	}

	return nil
}

// テーブルを削除します。
func (d *Database) DropTables() error {
	// users
	_, err := d.db.Exec("drop table users;")
	if err != nil {
		return err
	}

	// posts
	_, err = d.db.Exec("drop table posts;")
	if err != nil {
		return err
	}

	return nil
}
