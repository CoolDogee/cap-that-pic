package models

type Tag struct {
	Name       string
	Confidence float64
}

type TagList struct {
	List []Tag `json:"tags"`
}
