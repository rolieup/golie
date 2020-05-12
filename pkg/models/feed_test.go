package models_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rolieup/golie/pkg/models"
)

func TestParser_Parse(t *testing.T) {
	xmlFile := "../../examples/rolie/feed.xml"
	base := filepath.Base(xmlFile)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	fmt.Printf("Testing parsing of %s... ", name)
	bytes, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	err = xml.Unmarshal(bytes, &(models.Feed{}))
	if err != nil {
		panic(err)
	}
}
