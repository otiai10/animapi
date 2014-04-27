package anime

import "animapi/domain/song"
import "animapi/domain/person"
import "animapi/infrastructure/syobocal"

type AnimeFactory struct{}

func NewFactory() AnimeFactory {
	return AnimeFactory{}
}

func (f AnimeFactory) Dummy() Anime {
	return Anime{
		1,
		"凪のあすから",
		"コメントー",
		[]song.Song{
			song.Song{"edd and flow", 0, person.New("Ray")},
		},
		BroadcastPeriod{2013, 10},
		BroadcastPeriod{2014, 4},
	}
}

func (f AnimeFactory) Create(
	tid int,
	title string,
	comment string,
	songs []song.Song,
	start BroadcastPeriod,
	end BroadcastPeriod,
) Anime {
	return Anime{
		tid,
		title,
		comment,
		songs,
		start,
		end,
	}
}

func (f AnimeFactory) CreateListFromResponse(
	response syobocal.TitleLookupResponse,
) []Anime {
	return []Anime{f.Dummy()}
}
