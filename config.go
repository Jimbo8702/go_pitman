package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Config struct {
    MaxURLsToCrawl         int           `json:"max_urls_to_crawl"`
    CrawlTimeoutSeconds    time.Duration `json:"crawl_timeout_seconds"`
    MaxRequestsPerSecond   int           `json:"max_requests_per_second"`
    StartURL               string        `json:"start_url"`
    OutputFolder           string        `json:"output_folder"`
    OutputName             string        `json:"output_name"`
    OutputFileExtension    string        `json:"output_file_extension"`
    UserAgents             []string      `json:"user_agents"`
    IPAddress              []string      `json:"ip_addresses"`
}

func NewConfig(configFilePath string) (*Config, error) {
    configData, err := ioutil.ReadFile(configFilePath)
    if err != nil {
        return nil, fmt.Errorf("error reading the config file: %s", err)
    }

    var config Config
    if err := json.Unmarshal(configData, &config); err != nil {
        return nil, fmt.Errorf("error parsing the config file: %s", err)
    }

    return &config, nil
}