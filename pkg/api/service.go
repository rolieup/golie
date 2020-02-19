/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package api

import (
	"encoding/xml"
)

type JSONServiceRoot struct {
	Service Service `xml:"-" json:"service,"`
}

type Service struct {
	XMLName   xml.Name    `xml:"https://www.w3.org/2007/app service" json:"-"`
	Atomns    string      `xml:"xmlns:atom,attr" json:"-"`
	Workspace []Workspace `xml:"workspace" json:"workspace"`
}

type Workspace struct {
	Title      string       `xml:"atom:title" json:"title"`
	Collection []Collection `xml:"collection" json:"collection"`
}

type Collection struct {
	Href       string      `xml:"href,attr" json:"href"`
	Title      string      `xml:"atom:title" json:"title"`
	Link       *AtomLink   `xml:"atom:link" json:"link"`
	Categories *Categories `xml:"categories" json:"categories"`
}

type AtomLink struct {
	Link
}

type Categories struct {
	Fixed    string         `xml:"fixed,attr,omitempty" json:"fixed,omitempty"`
	Category []AtomCategory `xml:"atom:category" json:"category"`
}

type AtomCategory struct {
	Category
}
