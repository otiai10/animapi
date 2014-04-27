package song

import "animapi/domain/person"

var Type_OP = 0
var Type_ED = 1

type SongFactory struct{}

func NewFactory() SongFactory {
	return SongFactory{}
}

func (f SongFactory) Create(
	title string,
	songType int,
	singer person.Person,
) Song {
	return Song{
		title,
		songType,
		singer,
	}
}
