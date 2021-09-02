package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
	"github.com/urfave/cli/v2"
)

func generateAction(c *cli.Context) error {
	title := c.String("title")
	description := c.String("description")
	link := c.String("link")
	files := c.String("files")
	image := c.String("image")
	outFile := c.String("out-file")

	beginningOfTime := time.UnixMilli(0)

	// ToDo: validation

	fileInfos, err := ioutil.ReadDir(files)
	if err != nil {
		return err
	}

	feed := &feeds.Feed{
		Title:       title,
		Link:        &feeds.Link{Href: link},
		Description: description,
		Image: &feeds.Image{
			Url:   link + "/" + image,
			Title: image,
			Link:  link + "/" + image,
		},
		Created: beginningOfTime,
	}

	for i, fileInfo := range fileInfos {
		itemLink := link + "/" + fileInfo.Name()

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

func main() {
	app := &cli.App{
		Name:  "lipsync",
		Usage: "create podcast RSS files on the fly!",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generates a podcast RSS file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Usage:    "the title of the podcast",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "description",
						Usage: "the description of the podcast",
						Value: "",
					},
					&cli.StringFlag{
						Name:     "link",
						Usage:    "the base URL where the podcast will be served",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "files",
						Usage:    "the directory of the podcast files",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "image",
						Usage: "the path to the podcast image",
						Value: "",
					},
					&cli.StringFlag{
						Name:  "out-file",
						Usage: "the name of the output file",
						Value: "out.rss",
					},
				},
				Action: generateAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
