package blawg

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNewRSS(t *testing.T) {
	posts := Posts{
		Post{
			TitleText: "a bum title",
			Metadata: Metadata{
				Categories:  []string{"unix", "fun"},
				Description: "The best post ever",
			},
		},
	}

	config := Config{
		URL:         "http://test-url",
		Description: "an amazing site",
		Title:       "Test URL",
	}

	rss, err := newRSS(&posts, config)
	if err != nil {
		t.Errorf("Couldn't even make a new RSS")
	}
	t.Run("RSS uses the config description", func(t *testing.T) {
		if rss.Channel.Description != config.Description {
			t.Errorf("Expected %+s to be %+s\n", rss.Channel.Description, config.Description)
		}
	})

	t.Run("RSS uses the config title", func(t *testing.T) {
		if rss.Channel.Title != config.Title {
			t.Errorf("Expected %+s to be %+s\n", rss.Channel.Description, config.Description)
		}
	})

	t.Run("creates the expected item", func(t *testing.T) {
		want := RSSItem{
			Category:    strings.Join(posts[0].Metadata.Categories, ","),
			Description: posts[0].Metadata.Description,
			Title:       posts[0].TitleText,
			Link:        "http://test-urlposts/1/1/1/a-bum-title/",
			Guid:        "http://test-urlposts/1/1/1/a-bum-title/",
			PubDate:     "Mon, 01 Jan 0001 00:00:00 +0000",
		}
		got := rss.Channel.Item[0]
		if got != want {
			t.Errorf("wanted\n\t%#v\nbut got\n\t%#v\n", want, got)
		}
	})

	t.Run("RSS actually makes some RSS", func(t *testing.T) {
		sb := strings.Builder{}
		xml.NewEncoder(&sb).Encode(rss)
		rssXML := sb.String()
		want := `<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">`
		if !strings.Contains(rssXML, want) {
			t.Errorf("Expected %+s to contain %+s\n", rssXML, want)
		}
	})
}
