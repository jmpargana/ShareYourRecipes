package db

import (
	"context"
	"log"
	"server/models"
	"testing"
	"time"
)

func TestTagInsertAndFind(t *testing.T) {
	t.Parallel()

	// Needs to have tags not used by other tests
	tt := map[string]struct {
		tag      models.Tag
		id1, id2 int
	}{
		"pasta with two high ids": {"pasta", 5, 6},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assertNil(t, w.InsertTag(tc.id1, tc.tag))

			id1, err := w.FindTag(tc.tag)
			assertNil(t, err)

			// Insert Same Tag again and check
			assertNil(t, w.InsertTag(tc.id2, tc.tag))

			id2, err := w.FindTag(tc.tag)
			assertNil(t, err)

			if id1 != id2 {
				log.Fatalf("should receive same ids for same tag always, instead got: %d and %d", id1, id2)
			}

			// Check if two recipe_tags where created
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// stmt, err := w.db.PrepareContext(ctx, "SELECT recipeid FROM recipetags WHERE tagid = ?")
			// if err != nil {
			// 	log.Fatal(err)
			// }

			rows, err := w.db.QueryContext(ctx, "SELECT recipeid FROM recipetags WHERE tagid = ?", id2)
			if err != nil {
				log.Fatal(err)
			}

			counter := 0
			var ids []int
			for rows.Next() {
				var recipeid int
				rows.Scan(&recipeid)
				counter++
				ids = append(ids, recipeid)
			}

			if counter != 2 {
				t.Fatalf("should receive 2, instead got: %d", counter)
			}

			if ids[0] != tc.id1 || ids[1] != tc.id2 {
				t.Fatalf("got: %v, should get %d and %d", ids, tc.id1, tc.id2)
			}
		})
	}
}

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

			assertNil(t, w.InsertTags(1, tc.tags))

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
