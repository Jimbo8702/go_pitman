package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Downloader struct {
	Client   		 *http.Client
	UserAgent 			   string
	// IPAddress 			   string
	OutputFolder           string     
    OutputName             string      
    OutputFileExtension    string   
}

func NewDownloader(outFolder, outputName, outputFile string) *Downloader {
	return &Downloader{
		Client:   &http.Client{Timeout: 5 * time.Second},
		UserAgent: "",
		// IPAddress: "",
		OutputFolder: outFolder,
		OutputName: outputName,
		OutputFileExtension: outputFile,
	}
}

// func (d *Downloader) SetIPAddress(ipAddress string) {
// 	d.ipAddress = ipAddress
// }

func (d *Downloader) SetUserAgent(userAgent string) {
	d.UserAgent = userAgent 
}

func (d *Downloader) Download(url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the user agent in the request
	if d.UserAgent != "" {
		request.Header.Set("User-Agent", d.UserAgent)
	}

	//set the IP address for request
	// if d.ipAddress != "" {
	// 	request.Header.Set("X-Forwarded-For", d.ipAddress)
	// }

	response, err := d.Client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (d *Downloader) WriteDataToJSON(data ParsedData, urlNumber int) error {
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