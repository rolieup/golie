package utils

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/rolieup/golie/version"
	log "github.com/sirupsen/logrus"
)

func Acquire(URI string) (io.ReadCloser, error) {
	log.Infof("Downloading %s\n", URI)

	client := &http.Client{}
	// Make a request
	req, err := http.NewRequest("GET", URI, nil)
	if err != nil {
		return nil, err
	}
	// Send GolieVersion in Header
	useragent := fmt.Sprintf("Golie/%s (%s/%s)", version.Version, runtime.GOOS, runtime.GOARCH)
	req.Header.Add("User-Agent", useragent)
	req.Header.Set("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response %d from server on %s", response.StatusCode, URI)
	}
	return response.Body, nil
}
