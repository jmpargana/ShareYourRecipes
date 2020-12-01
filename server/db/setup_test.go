package db

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// func TestCreate(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "INSERT INTO users \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.Phone).WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := repo.Create(u)
// 	assert.NoError(t, err)
// }

// func TestCreateError(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "INSERT INTO user \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.Phone).WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := repo.Create(u)
// 	assert.Error(t, err)
// }

// func TestUpdate(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "UPDATE users SET name = \\?, email = \\?, phone = \\? WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.Name, u.Email, u.Phone, u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := repo.Update(u)
// 	assert.NoError(t, err)
// }

// func TestUpdateErr(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "UPDATE user SET name = \\?, email = \\?, phone = \\? WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.Name, u.Email, u.Phone, u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := repo.Update(u)
// 	assert.Error(t, err)
// }

// func TestDelete(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "DELETE FROM users WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := repo.Delete(u.ID)
// 	assert.NoError(t, err)
// }

// func TestDeleteError(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &DBWrapper{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "DELETE FROM user WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := repo.Delete(u.ID)
// 	assert.Error(t, err)
// }
