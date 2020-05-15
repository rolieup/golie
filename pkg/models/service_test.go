package models_test

import (
	"testing"

	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/stretchr/testify/assert"
)

func TestRFCServiceParsing(t *testing.T) {
	doc, err := rolie_source.ReadDocumentFromFile("../../examples/rolie/service/pcs.xml")
	assert.Nil(t, err)
	workspaces := doc.Service.Workspaces
	assert.Equal(t, len(workspaces), 2)
	assert.Equal(t, workspaces[0].Title, "Public Security Information Sharing")
	cols := workspaces[0].Collections
	assert.Equal(t, len(cols), 1)
	assert.Equal(t, cols[0].Href, "https://example.org/provider/public/vulns")
	assert.Equal(t, cols[0].Title, "Public Vulnerabilities")
	assert.Equal(t, cols[0].Link.Link.Href, "https://example.org/rolie/servicedocument")
	assert.Equal(t, cols[0].Link.Link.Rel, "service")
	assert.Equal(t, cols[0].Link.Link.Type, "")
	assert.Equal(t, cols[0].Link.Link.HrefLang, "")
	assert.Equal(t, cols[0].Link.Link.Title, "")
	assert.Equal(t, cols[0].Link.Link.Length, uint(0))
	assert.Equal(t, cols[0].Link.Link.Body, "")
	assert.Equal(t, cols[0].Categories.Fixed, "yes")
	categories := cols[0].Categories.Category
	assert.Equal(t, len(categories), 1)
	assert.Equal(t, categories[0].Category.Scheme, "urn:ietf:params:rolie:category:information-type")
	assert.Equal(t, categories[0].Category.Term, "vulnerability")

	assert.Equal(t, workspaces[1].Title, "Private Consortium Sharing")
	cols = workspaces[1].Collections
	assert.Equal(t, len(cols), 1)
	assert.Equal(t, cols[0].Href, "https://example.org/provider/private/incidents")
	assert.Equal(t, cols[0].Title, "Incidents")
	assert.Equal(t, cols[0].Link.Link.Href, "https://example.org/rolie/servicedocument")
	assert.Equal(t, cols[0].Link.Link.Rel, "service")
	assert.Equal(t, cols[0].Link.Link.Type, "")
	assert.Equal(t, cols[0].Link.Link.HrefLang, "")
	assert.Equal(t, cols[0].Link.Link.Title, "")
	assert.Equal(t, cols[0].Link.Link.Length, uint(0))
	assert.Equal(t, cols[0].Link.Link.Body, "")
	assert.Equal(t, cols[0].Categories.Fixed, "yes")
	categories = cols[0].Categories.Category
	assert.Equal(t, len(categories), 1)
	assert.Equal(t, categories[0].Category.Scheme, "urn:ietf:params:rolie:category:information-type")
	assert.Equal(t, categories[0].Category.Term, "incident")
}
