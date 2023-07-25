package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Downloader struct {
	Client   *http.Client
	OutputFolder           string     
    OutputName             string      
    OutputFileExtension    string   
}

func NewDownloader(outFolder, outputName, outputFile string) *Downloader {
	return &Downloader{
		Client:   &http.Client{Timeout: 5 * time.Second},
		OutputFolder: outFolder,
		OutputName: outputName,
		OutputFileExtension: outputFile,
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

func (d *Downloader) WriteDataToJSON(data Parseable, urlNumber int) error {
	filename := fmt.Sprintf("%s/%s_%s_.%s", d.OutputFolder, d.OutputName ,fmt.Sprint(urlNumber), d.OutputFileExtension)

	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = ioutil.WriteFile(filename, jsonString, 0644)
	if err != nil {
		return err
	}

	return nil
}