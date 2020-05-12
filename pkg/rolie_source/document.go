package rolie_source

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/rolieup/golie/pkg/models"
)

const (
	feedRootElement = "feed"
	// TODO Entry, Service, ...
)

// Rolie Document. Either Feed, Entry or Service
type Document struct {
	XMLName xml.Name `json:"-"`
	*models.Feed
}

func ReadDocument(r io.Reader) (*Document, error) {
	rawBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	d := xml.NewDecoder(bytes.NewReader(rawBytes))
	for {
		token, err := d.Token()
		if err != nil || token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			switch startElement.Name.Local {
			case feedRootElement:
				var feed models.Feed
				if err := d.DecodeElement(&feed, &startElement); err != nil {
					return nil, err
				}
				return &Document{Feed: &feed}, nil
			}
		}
	}

	var jsonTemp map[string]json.RawMessage
	if err := json.Unmarshal(rawBytes, &jsonTemp); err == nil {
		for k, v := range jsonTemp {
			switch k {
			case feedRootElement:
				var feed models.Feed
				if err := json.Unmarshal(v, &feed); err != nil {
					return nil, err
				}
				return &Document{Feed: &feed}, nil
			}
		}
	}

	return nil, errors.New("Malformed rolie document. Must be XML or JSON.")
}

func ReadDocumentFromFile(path string) (*Document, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ReadDocument(reader)
}
