package db

import (
	"context"
	"log"
	"server/models"
	"testing"
	"time"
)

func TestTagInsertAndFind(t *testing.T) {
	var tag models.Tag = "curry"

	assertNil(t, w.InsertTag(0, tag))

	id, err := w.FindTag(tag)
	assertNil(t, err)

	if id != 1 {
		t.Fatalf("shoul have received first id, instead got: %d", id)
	}

	// Insert Same Tag again and check
	assertNil(t, w.InsertTag(1, tag))

	id, err = w.FindTag(tag)
	assertNil(t, err)

	if id != 1 {
		t.Fatalf("shoul have received first id, instead got: %d", id)
	}

	// Check if two recipe_tags where created
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, "SELECT recipeid FROM recipetags where tagid = ?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	for rows.Next() {
		var recipeid int
		rows.Scan(&recipeid)
		counter++
	}

	if counter != 2 {
		t.Fatal("should receive 2")
	}
}

func TestInsertionMultipleTagsWithDuplicates(t *testing.T) {

}
