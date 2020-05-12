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
	xmlFile := "../../examples/rolie/feed/gov.nist.nvd.cve.recent.xml"
	base := filepath.Base(xmlFile)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	fmt.Printf("Testing parsing of %s... ", name)
	xmlReader, err := os.Open(xmlFile)
	if err != nil {
		panic(err)
	}
	xmlParsed := &models.Feed{}
	decoder := xml.NewDecoder(xmlReader)
	err = decoder.Decode(xmlParsed)
	if err != nil {
		panic(err)
	}

	jsonFile := "../../examples/rolie/feed/gov.nist.nvd.cve.recent.json"
	jsonParsedRoot := &models.JSONFeedRoot{}
	jsonBytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonBytes, jsonParsedRoot)
	if err != nil {
		panic(err)
	}
	jsonParsed := &jsonParsedRoot.Feed
	jsonParsed.XMLName = xmlParsed.XMLName // Disregard that json parser did not populate XMLName
	assert.Equal(t, xmlParsed, jsonParsed, "XML parser returns different data than json equivalent")
}
