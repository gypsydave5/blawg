package blawg

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

const timeFormat = time.RFC1123Z

func makeRSSFeed(siteDirectory string, config Config, posts *Posts) error {
	path := fmt.Sprintf("%s/feeds", siteDirectory)
	err := os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/feed.rss", path)
	f, err := os.Create(fileName)
	rss, _ := newRSS(posts, config)

	err = xml.NewEncoder(f).Encode(rss)
	return err
}

func newRSS(posts *Posts, config Config) (*RSS, error) {
	var rss RSS
	rss.Version = "2.0"
	rss.XMLName = xml.Name{Local: "rss", Space: "rss"}
	rss.Atom = "http://www.w3.org/2005/Atom"
	rss.Channel.AtomLink.Href = config.URL + "feeds/feed.rss"
	rss.Channel.AtomLink.Rel = "self"
	rss.Channel.AtomLink.XMLName = xml.Name{Local: "link", Space: "atom"}
	rss.Channel.AtomLink.Type = "application/rss+xml"
	rss.Channel.Link = config.URL + "feeds/feeds.rss"

	rss.Channel.Description = config.Description
	rss.Channel.PubDate = time.Now().Format(timeFormat)
	rss.Channel.LastBuildDate = time.Now().Format(timeFormat)
	rss.Channel.Title = config.Title
	rss.Channel.TTL = "1800"

	var items []RSSItem
	for i := 0; i < len([]Post(*posts)) && i < 10; i++ {
		var it RSSItem
		p := []Post(*posts)[i]
		it.Title = p.TitleText
		it.Description = p.Description
		it.Link = config.URL + "posts/" + p.Path()
		it.Guid = it.Link
		it.PubDate = p.Date.Format(timeFormat)
		it.Category = strings.Join(p.Categories, ",")
		it.Content = Content{Text: string(p.Body)}
		items = append(items, it)
	}

	rss.Channel.Item = items
	return &rss, nil
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"xmlns:atom,attr"`
	Channel struct {
		Text          string    `xml:",chardata"`
		Title         string    `xml:"title"`
		Description   string    `xml:"description"`
		LastBuildDate string    `xml:"lastBuildDate"`
		PubDate       string    `xml:"pubDate"`
		TTL           string    `xml:"ttl"`
		Link          string    `xml:"link"`
		AtomLink      AtomLink  `xml:"atom:link"`
		Item          []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string  `xml:"title"`
	Description string  `xml:"description"`
	Link        string  `xml:"link"`
	PubDate     string  `xml:"pubDate"`
	Category    string  `xml:"category"`
	Guid        string  `xml:"guid"`
	Content     Content `xml:"content"`
}

type Content struct {
	XMLName xml.Name `xml:"content"`
	Text    string   `xml:",cdata"`
}

type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	Text    string   `xml:",chardata"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}
