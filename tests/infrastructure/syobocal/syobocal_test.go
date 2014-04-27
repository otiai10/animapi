package syobocal

import "testing"
import . "github.com/r7kamura/gospel"
import "animapi/infrastructure/syobocal"

import "time"

// import "fmt"

func TestFindAfter(t *testing.T) {
	Describe(t, "FindAfter", func() {
		It("should find animes after param time", func() {
			client := syobocal.NewSyobocalClient()
			dur, _ := time.ParseDuration("-24h")
			_ = client.FindAfter(dur)
			// response := client.FindAfter(criteriaTime)
			// fmt.Printf("%+v", response)
		})
	})
}
