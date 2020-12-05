package models

import "sort"

type Recipe struct {
	ID          int
	UserID      int
	Title       string
	Ingridients []Ingridient
	Method      string
	Time        int
	Private     bool
	Tags        []Tag
}

func (r *Recipe) Equals(other *Recipe) bool {
	return r.ID == other.ID &&
		r.UserID == other.UserID &&
		r.Title == other.Title &&
		r.Method == other.Method &&
		r.Time == other.Time &&
		r.Private == other.Private &&
		r.compareIngridients(other) &&
		r.compareTags(other)
}

func (r *Recipe) compareIngridients(other *Recipe) bool {
	if len(r.Ingridients) != len(other.Ingridients) {
		return false
	}

	sort.Sort(ByIngridient(r.Ingridients))
	sort.Sort(ByIngridient(other.Ingridients))

	for i := range r.Ingridients {
		if r.Ingridients[i] != other.Ingridients[i] {
			return false
		}
	}

	return true
}

func (r *Recipe) compareTags(other *Recipe) bool {
	if len(r.Tags) != len(other.Tags) {
		return false
	}

	sort.Sort(ByTag(r.Tags))
	sort.Sort(ByTag(other.Tags))

	for i := range r.Tags {
		if r.Tags[i] != other.Tags[i] {
			return false
		}
	}

	return true
}
