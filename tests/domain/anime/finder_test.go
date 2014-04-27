package syobocal

import "testing"
import . "github.com/r7kamura/gospel"
import "animapi/domain/anime"

import "fmt"

func TestFinder(t *testing.T) {
	Describe(t, "FindByRange", func() {
		It("should find animes after param time", func() {
			finder := anime.NewFinder()

			animes := finder.FindByRange("-24h")

			fmt.Printf("Lengthほげ %v", len(animes))
		})
	})
}
