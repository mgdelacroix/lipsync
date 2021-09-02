package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

func main() {
	titleFlag := flag.String("title", "", "the title of the podcast")
	descriptionFlag := flag.String("description", "", "the description of the podcast")
	linkFlag := flag.String("link", "", "the link of the podcast")
	filesFlag := flag.String("files", ".", "the directory of the podcast files")
	outFile := flag.String("out-file", "out.rss", "the name of the output file")
	flag.Parse()

	beginningOfTime := time.UnixMilli(0)

	// ToDo: validation

	fileInfos, err := ioutil.ReadDir(*filesFlag)
	if err != nil {
		log.Fatal(err)
	}

	feed := &feeds.Feed{
		Title: *titleFlag,
		Link: &feeds.Link{Href: *linkFlag},
		Description: *descriptionFlag,
		Created: beginningOfTime,
	}

	for i, fileInfo := range fileInfos {
		itemLink := *linkFlag+"/"+fileInfo.Name()

		feed.Add(&feeds.Item{
			Title: fileInfo.Name(),
			Link: &feeds.Link{Href: itemLink},
			Description: fileInfo.Name(),
			Created: beginningOfTime.Add(time.Duration(i)*time.Second),
			Enclosure: &feeds.Enclosure{Url: itemLink, Length: strconv.FormatInt(fileInfo.Size(), 10), Type: "audio/mpeg"},
		})
	}

	rss, err := feed.ToRss()
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(*outFile, []byte(rss), 0755)
}
