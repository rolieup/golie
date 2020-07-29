package rolie_source

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/rolieup/golie/pkg/models"
)

const (
	feedRootElement    = "feed"
	entryRootElement   = "entry"
	serviceRootElement = "service"
)

// Rolie Document. Either Feed, Entry or Service
type Document struct {
	XMLName         xml.Name `json:"-"`
	*models.Feed    `json:"feed,omitempty"`
	*models.Entry   `json:"entry,omitempty"`
	*models.Service `json:"service,omitempty"`
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
				return &Document{Feed: &feed}, models.AssertAtomNamespace(feed.XMLName.Space)
			case entryRootElement:
				var entry models.Entry
				if err := d.DecodeElement(&entry, &startElement); err != nil {
					return nil, err
				}
				return &Document{Entry: &entry}, models.AssertAtomNamespace(entry.XMLName.Space)
			case serviceRootElement:
				var service models.Service
				if err := d.DecodeElement(&service, &startElement); err != nil {
					return nil, err
				}
				return &Document{Service: &service}, models.AssertAtomPublishingNamespace(service.XMLName.Space)
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
			case entryRootElement:
				var entry models.Entry
				if err := json.Unmarshal(v, &entry); err != nil {
					return nil, err
				}
				return &Document{Entry: &entry}, nil
			case serviceRootElement:
				var service models.Service
				if err := json.Unmarshal(v, &service); err != nil {
					return nil, err
				}
				return &Document{Service: &service}, nil
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

// Writes both json and xml files. Provide path without extension.
func (doc *Document) Write(filePath string) error {
	err := doc.WriteJSON(filePath + ".json")
	if err != nil {
		return err
	}

	return doc.WriteXML(filePath + ".xml")
}

func (doc *Document) WriteJSON(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return doc.JSON(file, true)
}

func (doc *Document) WriteXML(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return doc.XML(file, true)
}

// XML writes the Rolie object as XML to the given writer
func (doc *Document) XML(w io.Writer, prettify bool) error {
	w.Write([]byte(xml.Header))
	e := xml.NewEncoder(w)
	if prettify {
		e.Indent("", "  ")
	}
	return e.Encode(doc)
}

// JSON writes the Rolie object as JSON to the given writer
func (doc *Document) JSON(w io.Writer, prettify bool) error {
	e := json.NewEncoder(w)
	if prettify {
		e.SetIndent("", "  ")
	}

	return e.Encode(doc)
}

// MarshalXML marshals either a catalog or a profile
func (doc *Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if doc.Feed != nil {
		doc.Feed.Xmlns = models.Atom2005HttpsUri
		if err := e.Encode(doc.Feed); err != nil {
			return err
		}
	} else if doc.Entry != nil {
		doc.Entry.Xmlns = models.Atom2005HttpsUri
		if err := e.Encode(doc.Entry); err != nil {
			return err
		}
	} else if doc.Service != nil {
		doc.Service.Xmlns = models.AtomPublishingHttpsUri
		if err := e.Encode(doc.Service); err != nil {
			return err
		}
	} else {
		return errors.New("Cannot marshal empty rolie document")
	}
	return nil
}

func ensureXmlns(n *xml.Name, space, local string) {
	if n.Space == "" {
		n.Space = space
		n.Local = local
	}
	fmt.Println(n.Space)
}
