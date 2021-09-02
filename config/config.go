package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	PodcastURL  string `yaml:"podcastURL"`
	Directory   string `yaml:"directory"`
	Image       string `yaml:"image"`
}

func ReadConfig(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		return nil, fmt.Errorf("cannot parse config file: %w", err)
	}

	return &config, err
}
