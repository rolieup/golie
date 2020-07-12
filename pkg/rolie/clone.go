package rolie

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rolieup/golie/pkg/models"
	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/rolieup/golie/version"
)

func Clone(URI string, dir string) error {
	f := fetcher{
		URI:           URI,
		DirectoryPath: dir,
	}
	f.Init()
	return f.Clone()
}

type fetcher struct {
	URI           string
	BaseURI       string
	DirectoryPath string
}

func (f *fetcher) Init() {
	idx := strings.LastIndex(f.URI, "/")
	if idx != -1 && idx != len(f.URI) {
		f.BaseURI = f.URI[:idx]
	}
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
	return f.cloneFeed(document.Feed)
}

func (f *fetcher) cloneFeed(feed *models.Feed) error {
	for _, entry := range feed.Entry {
		if len(entry.Link) > 0 {
			err := f.storeRemoteResource(entry.Link[0].Href)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *fetcher) storeRemoteResource(URI string) error {
	mainResource, err := f.getRemoteResourceRaw(URI)
	if err != nil {
		return err
	}
	defer mainResource.Close()

	rawBytes, err := ioutil.ReadAll(mainResource)
	if err != nil {
		return err
	}
	return f.storeLocally(URI, rawBytes)
}

func (f *fetcher) getRemoteResourceRaw(URI string) (io.ReadCloser, error) {
	fmt.Printf("Downloading %s\n", URI)

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

func (f *fetcher) storeLocally(URI string, content []byte) error {
	path, err := f.filepath(URI)
	if err != nil {
		return err
	}

	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, os.FileMode(0755))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, content, 0644)
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
	if strings.HasPrefix(URI, f.BaseURI) {
		return strings.TrimPrefix(URI, f.BaseURI), nil
	}
	location, err := url.Parse(URI)
	if err != nil {
		return "", fmt.Errorf("Could not parse URL: %v %s", err, URI)
	}
	return filepath.Join(location.Hostname(), location.EscapedPath()), nil
}
