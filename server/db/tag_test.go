package db

import (
	"server/models"
	"testing"
)

func TestTagInsertAndFind(t *testing.T) {
	var tag models.Tag = "curry"

	if err := w.InsertTag(0, tag); err != nil {
		t.Fatal(err)
	}

	id, err := w.FindTag(tag)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Fatalf("shoul have received first id, instead got: %d", id)
	}
}
