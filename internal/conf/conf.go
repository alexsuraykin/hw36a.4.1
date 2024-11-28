package conf

import (
	"encoding/json"
	"os"
)

type BDConfig struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Table    string `json:"table:"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewBD(filePath string) (*BDConfig, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	bd := BDConfig{}
	err = json.Unmarshal(configFile, &bd)
	if err != nil {
		return nil, err
	}
	return &bd, nil
}

type RSSConfig struct {
	UrlsRSS       []string `json:"rss"`
	RequestPeriod int      `json:"request_period"`
}

func NewRSS(filePath string) (*RSSConfig, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	rss := RSSConfig{}
	err = json.Unmarshal(configFile, &rss)
	if err != nil {
		return nil, err
	}
	return &rss, nil
}
