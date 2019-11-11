package data

import (
	"encoding/json"

	"github.com/cooldogee/cap-that-pic/models"
	"github.com/gobuffalo/packr/v2"
)

var (
	song    models.SongList
	tag     models.TagList
	boxSong = packr.New("lyrics", "../../lyrics/")
	boxTag  = packr.New("tags", "../../tags/")

//	client  *mongo.Client
)

func loadFile(name string, box *packr.Box) []byte {
	res, err := box.Find(name)
	if err != nil {
		panic(err)
	}
	return res
}

func init() {
	Reload()
}

func Reload() {
	//	client = db.ConnectToDB()
	//	db.CloseConnectionDB(client)

	song = *new(models.SongList)
	tag = *new(models.TagList)
	json.Unmarshal(loadFile("lyrics.json", boxSong), &song)
	json.Unmarshal(loadFile("tags1.json", boxTag), &tag)
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

func Tag(limit int, offset int) *models.TagList {
	if limit == -1 {
		return &tag
	} else {
		var tmp models.TagList
		end := offset + limit + 1
		if offset >= len(tag.List) {
			return nil
		}
		if end > len(tag.List) {
			end = len(tag.List)
		}
		tmp.List = tag.List[offset : offset+limit]
		return &tmp
	}
}
