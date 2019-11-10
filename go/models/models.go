package models

// Song represemts our Song Object stored in DB
type Song struct {
	Year     int
	Chart    string
	ChartURL string
	Rank     int
	Title    string
	Artist   string
	Lyrics   string
}
