/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package models

import "encoding/xml"

type Service struct {
	XMLName    xml.Name    `xml:"service" json:"-"`
	Xmlns      string      `xml:"xmlns atom,attr" json:"-"`
	Lang       string      `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty" json:"-"`
	Workspaces []Workspace `xml:"workspace,omitempty" json:"workspace"`
}

type Workspace struct {
	Title       string       `xml:"title" json:"title"`
	Collections []Collection `xml:"collection,omitempty" json:"collection"`
}

type Collection struct {
	Href       string      `xml:"href,attr" json:"href"`
	Title      string      `xml:"title,omitempty" json:"title"`
	Link       *AtomLink   `xml:"link" json:"link"`
	Categories *Categories `xml:"categories,omitempty" json:"categories"`
}

type AtomLink struct {
	Link
}

type Categories struct {
	Fixed    string     `xml:"fixed,attr,omitempty" json:"fixed,omitempty"`
	Category []Category `xml:"category,omitempty" json:"category"`
}
