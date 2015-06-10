package newsexport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var needle = `<?xml version="1.0" encoding="UTF-8"?>
 <rss xmlns:yandex="http://news.yandex.ru" xmlns:media="http://search.yahoo.com/mrss/" xmlns:version="2.0">
  <item>
   <title>hello</title>
   <link></link>
   <pubDate></pubDate>
   <yandex:full-text>fulltext</yandex:full-text>
  </item>
  <item>
   <title>hello1</title>
   <link></link>
   <pubDate></pubDate>
   <yandex:full-text>fulltext1</yandex:full-text>
  </item>
  <item>
   <title>hello2</title>
   <link></link>
   <pubDate></pubDate>
   <yandex:full-text>fulltext2</yandex:full-text>
  </item>
 </rss>`

func TestMarshal(t *testing.T) {
	feed := &Feed{
		Yandex:  "http://news.yandex.ru",
		Media:   "http://search.yahoo.com/mrss/",
		Version: "2.0",
		Items: []Item{
			Item{Title: "hello", FullText: "fulltext"},
			Item{Title: "hello1", FullText: "fulltext1"},
			Item{Title: "hello2", FullText: "fulltext2"},
		},
	}
	bts, err := feed.Marshal()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, needle, string(bts), "sdfsdf")
}
