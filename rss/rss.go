package rss

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/mgdelacroix/lipsync/config"

	"github.com/gorilla/feeds"
)

func GeneratePodcastRss(cfg *config.Config) (string, error) {
	beginningOfTime := time.UnixMilli(0)

	fileInfos, err := ioutil.ReadDir(cfg.Directory)
	if err != nil {
		return "", err
	}

	feed := &feeds.Feed{
		Title:       cfg.Title,
		Link:        &feeds.Link{Href: cfg.PodcastURL},
		Description: cfg.Description,
		Image: &feeds.Image{
			Url:   cfg.PodcastURL + "/image",
			Title: cfg.Image,
			Link:  cfg.PodcastURL + "/image",
		},
		Created: beginningOfTime,
	}

	for i, fileInfo := range fileInfos {
		itemLink := cfg.PodcastURL + "/files/" + fileInfo.Name()

		feed.Add(&feeds.Item{
			Title:       fileInfo.Name(),
			Link:        &feeds.Link{Href: itemLink},
			Description: fileInfo.Name(),
			Created:     beginningOfTime.Add(time.Duration(i) * time.Second),
			Enclosure:   &feeds.Enclosure{Url: itemLink, Length: strconv.FormatInt(fileInfo.Size(), 10), Type: "audio/mpeg"},
		})
	}

	return feed.ToRss()
}
