package models

type Recipe struct {
	ID          int
	Title       string
	Ingridients string
	Method      string
	Time        int
	Private     bool
	Tags        []string
}
