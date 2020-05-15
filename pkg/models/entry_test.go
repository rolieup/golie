package models_test

import (
	"testing"

	"github.com/rolieup/golie/pkg/models"
	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/stretchr/testify/assert"
)

func TestRFCEntryParsing(t *testing.T) {
	doc, err := rolie_source.ReadDocumentFromFile("../../examples/rolie/entry/f63aafa9-4082-48a3-9ce6-97a2d69d4a9b.xml")
	assert.Nil(t, err)
	entry := doc.Entry
	assert.Equal(t, entry.XMLName.Space, "https://www.w3.org/2005/Atom")
	assert.Equal(t, entry.XMLName.Local, "entry")
	assert.Equal(t, entry.ID, "f63aafa9-4082-48a3-9ce6-97a2d69d4a9b")
	assert.Equal(t, entry.Title, "Sample Vulnerability")
	assert.Nil(t, entry.Link)
	assert.Equal(t, entry.Published, models.TimeStr("2015-08-04T18:13:51.0Z"))
	assert.Equal(t, entry.Updated, models.TimeStr("2015-08-05T18:13:51.0Z"))
	assert.Equal(t, entry.Title, "Sample Vulnerability")
	assert.Nil(t, entry.Author)
	assert.Equal(t, entry.Summary.Type, "")
	assert.Equal(t, entry.Summary.Src, "")
	assert.Equal(t, entry.Summary.Body, "A vulnerability issue identified by CVE-...")
	assert.Equal(t, entry.Format.Ns, "urn:ietf:params:xml:ns:exampleformat")
	assert.Equal(t, entry.Format.Version, "")
	assert.Nil(t, entry.Property)
}
