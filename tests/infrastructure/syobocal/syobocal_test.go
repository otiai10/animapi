package syobocal

import "testing"
import . "github.com/r7kamura/gospel"
import "animapi/infrastructure/syobocal"

import "time"

// import "fmt"

func TestFindAfter(t *testing.T) {
	Describe(t, "FindAfter", func() {
		It("should find animes after param time", func() {
			dur, _ := time.ParseDuration("-24h")
			criteriaTime := time.Now().Add(dur)
			client := syobocal.NewSyobocalClient()
			_ = client.FindAfter(criteriaTime)
			// response := client.FindAfter(criteriaTime)
			// fmt.Printf("%+v", response)
		})
	})
}
