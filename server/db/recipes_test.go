package db

// func TestRecipeInsertFind(t *testing.T) {
// 	t.Parallel()

// 	recipes := []*models.Recipe{
// 		{
// 			ID:          10,
// 			UserID:      2,
// 			Private:     true,
// 			Title:       "korma",
// 			Ingridients: []models.Ingridient{"rice", "tofu", "butter"},
// 			Time:        30,
// 			Method:      "cook",
// 			Tags:        []models.Tag{"indian", "curry"},
// 		},
// 	}

// 	for _, r := range recipes {
// 		assertNil(t, w.InsertRecipe(r))
// 		recipe, err := w.FindRecipeByID(r.ID)
// 		assertNil(t, err)

// 		if !recipe.Equals(r) {
// 			t.Fatalf("got: %v, expected: %v", recipe, r)
// 		}
// 	}
// }
