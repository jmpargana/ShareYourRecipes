package db

import "testing"

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func contains(s []int, i int) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
