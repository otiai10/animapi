package factory

import (
	"bytes"
	"regexp"

	"github.com/otiai10/animapi/models"
	"github.com/otiai10/animapi/syobocal"
)

var (
	songHeader = regexp.MustCompile("^\\*(オープニングテーマ|エンディングテーマ|挿入歌)([0-9]*)「(.+)」")
)

// ConvertTitleLookupResponseToAnime syobocalパッケージのTitleLookupResponseをanimapi/modelsパッケージのAnimeに変換します.
func ConvertTitleLookupResponseToAnime(tlr syobocal.TitleLookupResponse) ([]*models.Anime, error) {
	animes := []*models.Anime{}
	for _, item := range tlr.TitleItems.Items {
		anime := &models.Anime{
			ID:         item.TID,
			UpdatedAt:  item.LastUpdate.Value,
			Title:      item.Title,
			CommentRaw: item.Comment,
			Category:   models.Category(item.Category),
			Songs:      parseRawComment(item.TID, item.Comment),
		}
		animes = append(animes, anime)
	}
	return animes, nil
}

// parseRawComment Comment
// フィールドを1行1行解釈して、songsを作っていく
// 将来的にはsongs以外も返す？
func parseRawComment(animeID int, raw string) []models.Song {
	b := []byte(raw)
	rows := bytes.Split(b, []byte("\n"))
	songs := []models.Song{}
	f := false
	for _, row := range rows {
		if len(row) == 0 {
			f = false
			continue
		}
		matches := songHeader.FindAllSubmatch(row, -1)
		if len(matches) < 1 {
			if f {
				keyValue := bytes.Split(bytes.Trim(row, ":"), []byte(":"))
				if len(keyValue) < 2 {
					continue
				}
				// 最後の要素のAttributesにぶちこむ
				songs[len(songs)-1].Attributes[string(keyValue[0])] = string(keyValue[1])
			}
			continue
		}
		f = true
		song := models.Song{
			AnimeID:    animeID,
			Type:       string(matches[0][1]),
			Seq:        string(matches[0][2]),
			Title:      string(matches[0][3]),
			Attributes: map[string]string{},
		}
		songs = append(songs, song)
	}
	return songs
}
