package blawg

import (
	"fmt"
	"os"
)

func makeRSSFeed(siteDirectory string) error {
	path := fmt.Sprintf("%s/feeds", siteDirectory)
	err := os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s/feed.rss", path)
	_, err = os.Create(fileName)
	return err
}
