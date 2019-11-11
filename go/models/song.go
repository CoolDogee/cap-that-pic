package models

type Song struct {
	Year     int
	Chart    string
	ChartURL string
	Rank     int
	Song     string
	Artist   string
	Lyrics   string
}

type SongList struct {
	List []Song `json:"song"`
}
