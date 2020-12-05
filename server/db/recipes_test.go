package db

import (
	"server/models"
	"testing"
)

func TestRecipeInsertFind(t *testing.T) {
	t.Parallel()

	recipes := []*models.Recipe{
		{
			ID:          11,
			UserID:      7,
			Private:     true,
			Title:       "korma",
			Ingridients: []models.Ingridient{"rice", "tofu", "butter"},
			Time:        30,
			Method:      "cook",
			Tags:        []models.Tag{"indian", "curry"},
		},
		{
			ID:          1001,
			UserID:      7,
			Private:     true,
			Title:       "Chop Suoy",
			Ingridients: []models.Ingridient{"rice", "tofu", "noodles"},
			Time:        30,
			Method:      "cook",
			Tags:        []models.Tag{"chinese", "noodles", "spicy"},
		},
	}

	for _, r := range recipes {
		assertNil(t, w.InsertRecipe(r))
		recipe, err := w.FindRecipeByID(r.ID)
		assertNil(t, err)

		if !recipe.Equals(r) {
			t.Fatalf("got: %v, expected: %v", recipe, r)
		}
	}
}

func TestInsertAllRecipesAndFindOneByOne(t *testing.T) {
	recipes := []*models.Recipe{
		{
			ID:          11,
			UserID:      7,
			Private:     true,
			Title:       "korma",
			Ingridients: []models.Ingridient{"rice", "tofu", "butter"},
			Time:        30,
			Method:      "cook",
			Tags:        []models.Tag{"indian", "curry"},
		},
		{
			ID:          1002,
			UserID:      7,
			Private:     true,
			Title:       "Chop Suoy",
			Ingridients: []models.Ingridient{"rice", "tofu", "noodles"},
			Time:        30,
			Method:      "cook",
			Tags:        []models.Tag{"chinese", "noodles", "spicy"},
		},
	}

	for _, r := range recipes {
		assertNil(t, w.InsertRecipe(r))
	}

	for _, r := range recipes {
		recipe, err := w.FindRecipeByID(r.ID)
		assertNil(t, err)

		if !recipe.Equals(r) {
			t.Fatalf("got: %v, expected: %v", recipe, r)
		}
	}
}
