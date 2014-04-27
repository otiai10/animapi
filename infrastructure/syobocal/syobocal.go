package syobocal

// 基盤つくる？
import "net/http"
import "time"
import "io/ioutil"
import "encoding/xml"

// import "fmt"

type SyobocalHTTPClient struct {
	BaseURL string
}

// とりあえず
var baseURL = "http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20140426_000000-"

func NewSyobocalClient() SyobocalHTTPClient {
	return SyobocalHTTPClient{
		BaseURL: baseURL,
	}
}

func (c SyobocalHTTPClient) FindAfter(criteria time.Time) TitleLookupResponse {
	resp, _ := http.Get(c.BaseURL)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var syobocalTitleLookupResponse TitleLookupResponse
	_ = xml.Unmarshal(bodyBytes, &syobocalTitleLookupResponse)

	// fmt.Printf("%T\n", syobocalTitleLookupResponse.TitleItems.TitleItem[0])

	return syobocalTitleLookupResponse
}
