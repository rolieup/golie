/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package api

import (
	"encoding/xml"
)

type JSONEntryRoot struct {
	Entry AtomEntry `xml:"xml:-" json:"entry"`
}

type AtomEntry struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom entry" json:"-"`
	Lang    string   `xml:"xml:lang,attr" json:"-"`
	Roliens string   `xml:"xmlns:rolie,attr" json:"-"`
	Entry   *Entry   `xml:"entry" json:"entry"`
}
