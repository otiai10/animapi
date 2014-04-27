package song

import "animapi/domain/person"

type Song struct {
	Title    string
	SongType int
	Singer   person.Person
}
