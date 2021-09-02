package main

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/mgdelacroix/lipsync/config"

	"github.com/gorilla/feeds"
	"github.com/urfave/cli/v2"
)

func generateAction(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		return err
	}
	outFile := c.String("out-file")
	beginningOfTime := time.UnixMilli(0)

	fileInfos, err := ioutil.ReadDir(cfg.Directory)
	if err != nil {
		return err
	}

	feed := &feeds.Feed{
		Title:       cfg.Title,
		Link:        &feeds.Link{Href: cfg.PodcastURL},
		Description: cfg.Description,
		Image: &feeds.Image{
			Url:   cfg.PodcastURL + "/" + cfg.Image,
			Title: cfg.Image,
			Link:  cfg.PodcastURL + "/" + cfg.Image,
		},
		Created: beginningOfTime,
	}

	for i, fileInfo := range fileInfos {
		itemLink := cfg.PodcastURL + "/" + fileInfo.Name()

		feed.Add(&feeds.Item{
			Title:       fileInfo.Name(),
			Link:        &feeds.Link{Href: itemLink},
			Description: fileInfo.Name(),
			Created:     beginningOfTime.Add(time.Duration(i) * time.Second),
			Enclosure:   &feeds.Enclosure{Url: itemLink, Length: strconv.FormatInt(fileInfo.Size(), 10), Type: "audio/mpeg"},
		})
	}

	rss, err := feed.ToRss()
	if err != nil {
		return err
	}

	ioutil.WriteFile(outFile, []byte(rss), 0755)

	return nil
}
