package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// UserModel represent the user model
type UserModel struct {
	ID    string
	Name  string
	Email string
	Phone string
}

// repository represent the repository model
type DBWrapper struct {
	db *sql.DB
}

// NewRepository will create a variable that represent the Repository struct
func New(dialect, dsn string, idleConn, maxConn int) (*DBWrapper, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &DBWrapper{db}, nil
}

func (w *DBWrapper) Close() {
	w.db.Close()
}
