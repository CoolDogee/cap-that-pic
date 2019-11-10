package models

type Tag struct {
	Name       string
	confidence float64
}

type TagList struct {
	List []Tag `json:"tag"`
}
