package models

type Recipe struct {
	ID          int
	Title       string
	Ingridients string
	Method      string
	Time        int
	Private     int
	// Tags        []string
}

func (r *Recipe) Equals(other *Recipe) bool {
	return r.ID == other.ID &&
		r.Title == other.Title &&
		r.Ingridients == other.Ingridients &&
		r.Method == other.Method &&
		r.Time == other.Time &&
		r.Private == other.Private
}
