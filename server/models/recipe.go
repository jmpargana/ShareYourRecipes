package models

type Recipe struct {
	ID          int
	Title       string
	Ingridients []string
	Method      string
	Time        int
	Private     int
	Tags        []string
}

func (r *Recipe) Equals(other *Recipe) bool {
	return r.ID == other.ID &&
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

	for i := range r.Tags {
		if r.Tags[i] != other.Tags[i] {
			return false
		}
	}

	return true
}
