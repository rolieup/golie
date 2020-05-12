package models_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rolieup/golie/pkg/models"
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
		xmlReader, err := os.Open(xmlFile)
		assert.Nil(t, err)
		xmlParsed := &models.Feed{}
		decoder := xml.NewDecoder(xmlReader)
		assert.Nil(t, decoder.Decode(xmlParsed))

		jsonFile := fmt.Sprintf("../../examples/rolie/feed/%s.json", name)
		jsonParsedRoot := &models.JSONFeedRoot{}
		jsonBytes, err := ioutil.ReadFile(jsonFile)
		assert.Nil(t, err)

		assert.Nil(t, json.Unmarshal(jsonBytes, jsonParsedRoot))
		jsonParsed := &jsonParsedRoot.Feed
		jsonParsed.XMLName = xmlParsed.XMLName // Disregard that json parser did not populate XMLName
		assert.Equal(t, xmlParsed, jsonParsed, "XML parser returns different data than json equivalent")
	}
}
