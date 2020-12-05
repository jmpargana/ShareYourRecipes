package models

type Ingridient string

type ByIngridient []Ingridient

func (t ByIngridient) Len() int           { return len(t) }
func (t ByIngridient) Less(i, j int) bool { return t[i] < t[j] }
func (t ByIngridient) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
