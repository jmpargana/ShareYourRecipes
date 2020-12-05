package models

type Tag string

type ByTag []Tag

func (t ByTag) Len() int           { return len(t) }
func (t ByTag) Less(i, j int) bool { return t[i] < t[j] }
func (t ByTag) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
