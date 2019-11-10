package data

import (
	"encoding/json"

	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gobuffalo/packr/v2"
)

var (
	song models.SongList
	box  = packr.New("lyrics", "../../lyrics/")
)

func loadFile(name string) []byte {
	res, err := box.Find(name)
	if err != nil {
		panic(err)
	}
	return res
}

func Reload() {
	song = *new(models.SongList)
	json.Unmarshal(loadFile("lyrics.json"), &song)
}

func Song(limit int, offset int) *models.SongList {
	if limit == -1 {
		return &song
	} else {
		var tmp models.SongList
		end := offset + limit + 1
		if offset >= len(song.List) {
			return nil
		}
		if end > len(song.List) {
			end = len(song.List)
		}
		tmp.List = song.List[offset : offset+limit]
		return &tmp
	}
}
