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
	feedFiles, err := filepath.Glob("../../examples/rolie/*/*.xml")
	assert.Nil(t, err)
	assert.NotEmpty(t, feedFiles)
	for _, xmlFile := range feedFiles {
		dir := filepath.Dir(xmlFile)
		base := filepath.Base(xmlFile)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing parsing of %s/%s\n", dir, name)
		xmlDoc, err := rolie_source.ReadDocumentFromFile(xmlFile)
		assert.Nil(t, err)

		jsonFile := fmt.Sprintf("%s/%s.json", dir, name)
		jsonDoc, err := rolie_source.ReadDocumentFromFile(jsonFile)
		assert.Nil(t, err)

		// Disregard that json parser did not populate XMLName
		if jsonDoc.Feed != nil {
			jsonDoc.Feed.XMLName = xmlDoc.Feed.XMLName
			for idx, entry := range xmlDoc.Feed.Entry {
				jsonDoc.Feed.Entry[idx].XMLName = entry.XMLName
			}
		} else if jsonDoc.Entry != nil {
			jsonDoc.Entry.XMLName = xmlDoc.Entry.XMLName
		} else if jsonDoc.Service != nil {
			jsonDoc.Service.XMLName = xmlDoc.Service.XMLName
			jsonDoc.Service.Xmlns = xmlDoc.Service.Xmlns

		}
		assert.Equal(t, xmlDoc, jsonDoc, "XML parser returns different data than json equivalent")
	}
}
