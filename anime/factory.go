package anime

type AnimeFactory struct{}

func NewFactory() AnimeFactory {
	return AnimeFactory{}
}

func CreateAnime() Anime {
	return Anime{
		TID:     1,
		Title:   "凪のあすから",
		Comment: "ほげふが",
		Start: BroadcastPeriod{
			2013,
			10,
		},
		End: BroadcastPeriod{
			0,
			0,
		},
	}
}
