package models_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/stretchr/testify/assert"
)

func TestCompareXMLandJSON(t *testing.T) {
	feedFiles, err := filepath.Glob("../../examples/rolie/feed/*.xml")
	assert.Nil(t, err)
	assert.NotEmpty(t, feedFiles)
	for _, xmlFile := range feedFiles {
		base := filepath.Base(xmlFile)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing parsing of %s... ", name)
		xmlDoc, err := rolie_source.ReadDocumentFromFile(xmlFile)
		assert.Nil(t, err)

		jsonFile := fmt.Sprintf("../../examples/rolie/feed/%s.json", name)
		jsonDoc, err := rolie_source.ReadDocumentFromFile(jsonFile)
		assert.Nil(t, err)

		jsonDoc.Feed.XMLName = xmlDoc.Feed.XMLName // Disregard that json parser did not populate XMLName
		assert.Equal(t, xmlDoc, jsonDoc, "XML parser returns different data than json equivalent")
	}
}
