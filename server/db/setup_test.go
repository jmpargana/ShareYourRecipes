package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDBConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed opening stub database connection: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"col1", "col2"})
	rows.AddRow("val1", "val2")

	DB := &DB{db}
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	DB.SelectRecipeByID(2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("something failed: %s", err)
	}
}
