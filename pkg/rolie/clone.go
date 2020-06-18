package rolie

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/rolieup/golie/pkg/rolie_source"
)

func Clone(URI string, dir string) error {
	f := fetcher{
		URI:           URI,
		DirectoryPath: dir,
	}
	return f.Clone()
}

type fetcher struct {
	URI           string
	DirectoryPath string
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
	err = f.storeLocally(f.URI, rawBytes)
	if err != nil {
		return err
	}

	document, err := rolie_source.ReadDocument(bytes.NewReader(rawBytes))
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

func (f *fetcher) storeLocally(URI string, content []byte) error {
	filepath, err := f.filepath(f.URI)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, content, 0777)
}

func (f *fetcher) filepath(URI string) (string, error) {
	path, err := f.filepathRelative(URI)
	if err != nil {
		return "", err
	}
	return filepath.Join(f.DirectoryPath, path), nil
}

func (f *fetcher) filepathRelative(URI string) (string, error) {
	if URI == f.URI {
		idx := strings.LastIndex(URI, "/")
		if idx != -1 && idx != len(URI) {
			return URI[idx:], nil
		}
	}
	if strings.HasPrefix(URI, f.URI) {
		return strings.TrimPrefix(URI, f.URI), nil
	}
	return "", fmt.Errorf("Not implemented yet")

}
