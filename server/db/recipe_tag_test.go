package db

import "testing"

type recipetag struct {
	r int
	t int
}

func TestRecipeTagsInsertAndFind(t *testing.T) {
	t.Parallel()

	tt := map[string]struct {
		rt       []recipetag
		expected []int
	}{
		"no duplicates": {
			rt: []recipetag{
				{1, 2},
			},
			expected: []int{1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			for _, rt := range tc.rt {
				if err := w.InsertRecipeTag(rt.r, rt.t); err != nil {
					t.Fatalf("failed insert recipe tag with: %s", err)
				}
			}

			got, err := w.FindRecipeTagByTagID(tc.rt[0].t)
			if err != nil {
				t.Fatalf("something failed: %s", err)
			}

			if len(got) != len(tc.rt) {
				t.Fatalf("failed reading. got: %v", got)
			}
		})
	}
}
