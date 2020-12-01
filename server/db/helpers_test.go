package db

import (
	"server/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func setupMock() (*DBWrapper, sqlmock.Sqlmock, func()) {
	db, mock := NewMock()
	w := &DBWrapper{db}
	return w, mock, func() {
		w.Close()
	}
}

func compareRecipes(t *testing.T, lhs, rhs *models.Recipe) {
	if !lhs.Equals(rhs) {
		t.Fatalf("expected %v, got %v", rhs, lhs)
	}
}

func checkNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("shouldn't receive error, got: %s", err)
	}
}

func recipeNotNil(t *testing.T, recipe *models.Recipe) {
	if recipe != nil {
		t.Fatalf("shouldn't find anything, %v", recipe)
	}
}

func shouldntBeNil(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("should have failed!")
	}
}
