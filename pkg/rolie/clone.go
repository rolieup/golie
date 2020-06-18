package rolie

import (
	"bytes"
	"fmt"
	"github.com/rolieup/golie/pkg/rolie_source"
	"io"
	"io/ioutil"
	"net/http"
)

func Clone(URI string) error {
	f := fetcher{
		URI: URI,
	}
	return f.Clone()
}

type fetcher struct {
	URI string
}

func (f *fetcher) Clone() error {
	mainResource, err := f.getRemoteResourceRaw(f.URI)
	if err != nil {
		return err
	}
	defer mainResource.Close()

	rawBytes, err := ioutil.ReadAll(mainResource)
	if err != nil {
		return err
	}
	// TODO store to disk
	mainResourceCopy := bytes.NewReader(rawBytes)

	document, err := rolie_source.ReadDocument(mainResourceCopy)
	if err != nil {
		return fmt.Errorf("Failed to parse rolie document %s", err)
	}
	if document.Feed == nil {
		return fmt.Errorf("Not implemented yet: Found ROLIE resource that is not rolie:feed.")
	}
	// TODO process feed
	return nil
}

func (f *fetcher) getRemoteResourceRaw(URI string) (io.ReadCloser, error) {
	client := &http.Client{}
	// Make a request
	req, err := http.NewRequest("GET", URI, nil)
	if err != nil {
		return nil, err
	}
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
