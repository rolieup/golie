package rolie

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/rolieup/golie/pkg/models"
	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/rolieup/golie/pkg/utils"
	log "github.com/sirupsen/logrus"
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
	mainResource, err := utils.Acquire(f.URI)
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
		err := f.cloneEntry(entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fetcher) cloneEntry(entry *models.Entry) error {
	if len(entry.Link) > 0 {
		if f.needsRefresh(entry) {
			return f.storeRemoteResource(entry.Link[0].Href)
		}
	}
	return nil
}

func (f *fetcher) storeRemoteResource(URI string) error {
	mainResource, err := utils.Acquire(URI)
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

func (f *fetcher) needsRefresh(entry *models.Entry) bool {
	link := entry.Link[0]
	if link.Length == 0 {
		return true
	}
	path, err := f.filepath(link.Href)
	if err != nil {
		log.Debugf("cannot compose local filepath for remote resource: %v", err)
		return true
	}
	stat, err := os.Stat(path)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		return true
	}
	if link.Length != uint64(stat.Size()) {
		return true
	}
	entryTime, err := entry.Updated.Time()
	if err != nil {
		return true
	}
	if stat.ModTime().Before(entryTime) {
		return true
	}

	log.Debugf("Skipping download of %s", link.Href)

	return false
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
