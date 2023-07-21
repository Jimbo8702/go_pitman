package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Downloader struct {
	Client   *http.Client
}

func NewDownloader() *Downloader {
	return &Downloader{
		Client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (d *Downloader) Download(url string) (string, error) {
	resp, err := d.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}