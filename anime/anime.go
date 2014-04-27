package anime

type Anime struct {
	TID     int
	Title   string
	Comment string
	Start   BroadcastPeriod
	End     BroadcastPeriod
}
type BroadcastPeriod struct {
	Year int
	Mont int
}
