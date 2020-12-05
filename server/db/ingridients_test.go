package db

import (
	"server/models"
	"testing"
)

// func TestIngridientInsertAndFind(t *testing.T) {
// 	t.Parallel()

// 	// Needs to have tags not used by other tests
// 	tt := map[string]struct {
// 		ingridient models.Ingridient
// 		id1, id2   int
// 	}{
// 		"pasta with two high ids": {"pasta", 10, 11},
// 	}

// 	for name, tc := range tt {
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			assertNil(t, w.InsertIngridient(tc.id1, tc.ingridient))

// 			id1, err := w.FindIngridient(tc.ingridient)
// 			assertNil(t, err)

// 			// Insert Same Tag again and check
// 			assertNil(t, w.InsertIngridient(tc.id2, tc.ingridient))

// 			id2, err := w.FindIngridient(tc.ingridient)
// 			assertNil(t, err)

// 			if id1 != id2 {
// 				log.Fatalf("should receive same ids for same ingridient always, instead got: %d and %d", id1, id2)
// 			}

// 			// Check if two recipe_tags where created
// 			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 			defer cancel()

// 			rows, err := w.db.QueryContext(ctx, "SELECT recipeid FROM recipeingridients WHERE ingridientid = ?", id2)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			counter := 0
// 			var ids []int
// 			for rows.Next() {
// 				var recipeid int
// 				rows.Scan(&recipeid)
// 				counter++
// 				ids = append(ids, recipeid)
// 			}

// 			if counter != 2 {
// 				t.Fatalf("should receive 2, instead got: %d, with: %v", counter, ids)
// 			}

// 			if ids[0] != tc.id1 || ids[1] != tc.id2 {
// 				t.Fatalf("got: %v, should get %d and %d", ids, tc.id1, tc.id2)
// 			}
// 		})
// 	}
// }

func TestInsertionMultipleIngridientsWithDuplicates(t *testing.T) {
	t.Parallel()

	tt := map[string]struct {
		ingridients []models.Ingridient
		dedup       int
	}{
		"unused tags in different tests": {
			ingridients: []models.Ingridient{"noodles", "asian", "chinese"},
			dedup:       3,
		},
		"tags used in other tests": {
			ingridients: []models.Ingridient{"curry", "indian", "korma"},
			dedup:       3,
		},
		"repeated tags": {
			ingridients: []models.Ingridient{"indian", "indian", "curry"},
			dedup:       2,
		},
		"repeated tags 2": {
			ingridients: []models.Ingridient{"curry", "indian", "indian", "curry", "asian", "asian", "asian"},
			dedup:       3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assertNil(t, w.InsertIngridients(1, tc.ingridients))

			var ids []int

			for _, ingridient := range tc.ingridients {
				id, err := w.FindIngridient(ingridient)
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

func TestFindIngridientByRecipeID(t *testing.T) {
	t.Parallel()

	recipe1 := &models.Recipe{
		ID:          30,
		UserID:      50,
		Private:     false,
		Title:       "cozido",
		Ingridients: []models.Ingridient{"arroz", "chourico"},
		Time:        30,
		Method:      "cozer",
		Tags:        []models.Tag{"portuguese"},
	}

	recipe2 := &models.Recipe{
		ID:          31,
		UserID:      50,
		Private:     false,
		Title:       "borrego",
		Ingridients: []models.Ingridient{"carne", "borrego", "arroz"},
		Time:        30,
		Method:      "cozer",
		Tags:        []models.Tag{"portuguese"},
	}

	assertNil(t, w.InsertRecipe(recipe1))
	assertNil(t, w.InsertRecipe(recipe2))

	recipe, err := w.FindRecipeByID(recipe1.ID)
	assertNil(t, err)

	if len(recipe.Ingridients) != len(recipe1.Ingridients) {
		t.Fatalf("failed, got: %v, expected: %v", recipe.Ingridients, recipe1.Ingridients)
	}
}
