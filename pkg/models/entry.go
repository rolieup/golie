/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package models

import (
	"encoding/xml"
)

type JSONEntryRoot struct {
	Entry AtomEntry `xml:"xml:-" json:"entry"`
}

type AtomEntry struct {
	XMLName xml.Name `xml:"entry" json:"-"`
	Lang    string   `xml:"xml:lang,attr" json:"-"`
	Roliens string   `xml:"xmlns:rolie,attr" json:"-"`
	Entry   *Entry   `xml:"entry" json:"entry"`
}
