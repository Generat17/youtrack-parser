package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AgileId           string               `yaml:"agile_id"`
	IsSkipDrafts      bool                 `yaml:"skip_drafts"`
	StatesWhitelist   []string             `yaml:"state_whitelist"`
	IssueFields       []string             `yaml:"issue_fields"`
	HistoryCategories []string             `yaml:"history_categories"`
	HistoryFields     []string             `yaml:"history_fields"`
	ListNormalNames   map[string][2]string `yaml:"normal_names"`
}

func NewConfig(path string) (Config, error) {
	return readConfigFile(path)
}

func readConfigFile(file string) (Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshalling config data: %v", err)
	}

	return config, nil
}
