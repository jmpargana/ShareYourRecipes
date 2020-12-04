package db

import (
	"server/models"
	"testing"
)

// func TestTagInsertAndFind(t *testing.T) {
// 	t.Parallel()

// 	var tag models.Tag = "curry"

// 	assertNil(t, w.InsertTag(0, tag))

// 	id, err := w.FindTag(tag)
// 	assertNil(t, err)

// 	if id != 1 {
// 		t.Fatalf("shoul have received first id, instead got: %d", id)
// 	}

// 	// Insert Same Tag again and check
// 	assertNil(t, w.InsertTag(1, tag))

// 	id, err = w.FindTag(tag)
// 	assertNil(t, err)

// 	if id != 1 {
// 		t.Fatalf("shoul have received first id, instead got: %d", id)
// 	}

// 	// Check if two recipe_tags where created
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	stmt, err := w.db.PrepareContext(ctx, "SELECT recipeid FROM recipetags where tagid = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := stmt.Query(1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	counter := 0
// 	for rows.Next() {
// 		var recipeid int
// 		rows.Scan(&recipeid)
// 		counter++
// 	}

// 	if counter != 2 {
// 		t.Fatal("should receive 2")
// 	}
// }

func TestInsertionMultipleTagsWithDuplicates(t *testing.T) {
	t.Parallel()

	tt := map[string]struct {
		tags  []models.Tag
		dedup int
	}{
		"unused tags in different tests": {
			tags:  []models.Tag{"noodles", "asian", "chinese"},
			dedup: 3,
		},
		"tags used in other tests": {
			tags:  []models.Tag{"curry", "indian", "korma"},
			dedup: 3,
		},
		"repeated tags": {
			tags:  []models.Tag{"indian", "indian", "curry"},
			dedup: 2,
		},
		"repeated tags 2": {
			tags:  []models.Tag{"curry", "indian", "indian", "curry", "asian", "asian", "asian"},
			dedup: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assertNil(t, w.InsertTags(0, tc.tags))

			var ids []int

			for _, tag := range tc.tags {
				id, err := w.FindTag(tag)
				assertNil(t, err)

				if !contains(ids, id) {
					ids = append(ids, id)
				}
			}

			if len(ids) != tc.dedup {
				t.Fatalf("shouldn find only 3 IDs, instead got: %d. %v", len(ids), ids)
			}

		})
	}
}

func TestInvalidTag(t *testing.T) {

}
