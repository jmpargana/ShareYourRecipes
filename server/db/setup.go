package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

	w := &DBWrapper{db}

	if err := w.CreateTables(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *DBWrapper) Close() {
	w.db.Close()
}

func (w *DBWrapper) CreateTables() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, q := range tables {
		_, err := w.db.ExecContext(ctx, q)

		if err != nil {
			return fmt.Errorf("failed creating DB with: %s", err)
		}
	}

	return nil
}

func GenerateDSN(user, password, endPoint, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, endPoint, database)
}
