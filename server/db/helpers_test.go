package db

import "testing"

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
