package models_test

import (
	"testing"

	"github.com/rolieup/golie/pkg/models"
	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := rolie_source.ReadDocumentFromFile("../../examples/rolie/feed/2a7e265a-39bc-43f2-b711-b8fd9264b5c9.xml")
	assert.Nil(t, err)
	feed := doc.Feed
	assert.Equal(t, feed.XMLName.Space, "https://www.w3.org/2005/Atom")
	assert.Equal(t, feed.XMLName.Local, "feed")
	assert.Equal(t, feed.Lang, "")
	assert.Equal(t, feed.Roliens, "")
	assert.Equal(t, feed.ID, "2a7e265a-39bc-43f2-b711-b8fd9264b5c9")
	assert.Equal(t, feed.Title, "\n      Atom-formatted representation of\n      a Feed of XML vulnerability documents\n  ")
	assert.Equal(t, feed.Description, "")
	assert.Nil(t, feed.Author)
	assert.Nil(t, feed.Contributor)
	assert.Nil(t, feed.Generator)
	assert.Equal(t, feed.Rights, "")
	links := feed.Link
	assert.Equal(t, len(links), 2)
	assert.Equal(t, links[0].Rel, "self")
	assert.Equal(t, links[0].Href, "https://example.org/provider/public/vulns")
	assert.Equal(t, links[1].Rel, "service")
	assert.Equal(t, links[1].Href, "https://example.org/rolie/servicedocument")
	assert.Equal(t, feed.Category.Scheme, "urn:ietf:params:rolie:category:information-type")
	assert.Equal(t, feed.Category.Term, "vulnerability")
	assert.Equal(t, feed.Updated, models.TimeStr("2016-05-04T18:13:51.0Z"))
	entries := feed.Entry
	assert.Equal(t, len(entries), 1)
	assert.Equal(t, entries[0].ID, "dd786dba-88e6-440b-9158-b8fae67ef67c")
	assert.Equal(t, entries[0].Title, "Sample Vulnerability")
	assert.Nil(t, entries[0].Link)
	assert.Equal(t, entries[0].Published, models.TimeStr("2015-08-04T18:13:51.0Z"))
	assert.Equal(t, entries[0].Updated, models.TimeStr("2015-08-05T18:13:51.0Z"))
	assert.Equal(t, entries[0].Title, "Sample Vulnerability")
	assert.Nil(t, entries[0].Author)
	assert.Equal(t, entries[0].Summary.Type, "")
	assert.Equal(t, entries[0].Summary.Src, "")
	assert.Equal(t, entries[0].Summary.Body, "A vulnerability issue identified by CVE-...")
	assert.Equal(t, entries[0].Format.Ns, "urn:ietf:params:xml:ns:exampleformat")
	assert.Equal(t, entries[0].Format.Version, "")
	assert.Nil(t, entries[0].Property)

}
