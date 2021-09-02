package commands

import (
	"io/ioutil"

	"github.com/mgdelacroix/lipsync/config"
	"github.com/mgdelacroix/lipsync/rss"

	"github.com/urfave/cli/v2"
)

func GenerateAction(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		return err
	}
	outFile := c.String("out-file")

	rssStr, err := rss.GeneratePodcastRss(cfg)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outFile, []byte(rssStr), 0755); err != nil {
		return err
	}

	return nil
}
