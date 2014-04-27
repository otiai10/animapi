package anime

import "animapi/domain/song"

type Anime struct {
	TID     int
	Title   string
	Comment string
	Songs   []song.Song
	Start   BroadcastPeriod
	End     BroadcastPeriod
}
type BroadcastPeriod struct {
	Year  int
	Month int
}
