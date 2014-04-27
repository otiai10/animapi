package anime

import "animapi/infrastructure/syobocal"
import "time"

type Finder struct{}

func NewFinder() Finder {
	return Finder{}
}

func (f Finder) FindByRange(hours string) (animes []Anime) {

	dur, _ := time.ParseDuration("-" + hours)

	client := syobocal.NewSyobocalClient()
	response := client.FindAfter(dur)

	factory := NewFactory()
	animes = factory.CreateListFromResponse(response)

	return
}
