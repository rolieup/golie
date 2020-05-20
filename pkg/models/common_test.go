package models_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rolieup/golie/pkg/rolie_source"
	"github.com/stretchr/testify/assert"
)

func TestImportXMLandExportJSON(t *testing.T) {
	xmlFiles, err := filepath.Glob("../../examples/rolie/*/*.xml")
	assert.Nil(t, err)
	assert.NotEmpty(t, xmlFiles)
	for _, xmlFile := range xmlFiles {
		dir := filepath.Dir(xmlFile)
		base := filepath.Base(xmlFile)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		xmlDoc, err := rolie_source.ReadDocumentFromFile(xmlFile)
		assert.Nil(t, err)

		fmt.Printf("Testing conversion of %s/%s.xml to json\n", dir, name)
		var testJson strings.Builder
		err = xmlDoc.JSON(&testJson, true)
		assert.Nil(t, err)

		jsonFile := fmt.Sprintf("%s/%s.json", dir, name)
		jsonDoc, err := ioutil.ReadFile(jsonFile)
		assert.Nil(t, err)

		assert.Equal(t, testJson.String(), msWindowsPleaseDitchCarriageReturn(jsonDoc))
	}
}

func TestCompareXMLandJSON(t *testing.T) {
	xmlFiles, err := filepath.Glob("../../examples/rolie/*/*.xml")
	assert.Nil(t, err)
	assert.NotEmpty(t, xmlFiles)
	for _, xmlFile := range xmlFiles {
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
			jsonDoc.Feed.Xmlns = xmlDoc.Feed.Xmlns
			for idx, entry := range xmlDoc.Feed.Entry {
				jsonDoc.Feed.Entry[idx].XMLName = entry.XMLName
			}
		} else if jsonDoc.Entry != nil {
			jsonDoc.Entry.XMLName = xmlDoc.Entry.XMLName
			jsonDoc.Entry.Xmlns = xmlDoc.Entry.Xmlns
		} else if jsonDoc.Service != nil {
			jsonDoc.Service.XMLName = xmlDoc.Service.XMLName
			jsonDoc.Service.Xmlns = xmlDoc.Service.Xmlns

		}
		assert.Equal(t, xmlDoc, jsonDoc, "XML parser returns different data than json equivalent")
	}
}

func msWindowsPleaseDitchCarriageReturn(in []byte) string {
	return strings.ReplaceAll(string(in), "\r", "")
}
