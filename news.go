package newsexport

import (
	"encoding/xml"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Yandex  string   `xml:"xmlns:yandex,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Version string   `xml:"xmlns:version,attr"`
	Items   []Item   `xml:"channel"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	PubDate string   `xml:"pubDate"`

	FullText string `xml:"yandex:full-text"`
}

func NewFeed() *Feed {
	return &Feed{
		Yandex:  "http://news.yandex.ru",
		Media:   "http://search.yahoo.com/mrss/",
		Version: "2.0",
	}
}

func (f *Feed) Marshal() ([]byte, error) {
	data, err := xml.MarshalIndent(f, " ", " ") //xml.Marshal(f)
	data = append([]byte(xml.Header), data...)
	return data, err
}
