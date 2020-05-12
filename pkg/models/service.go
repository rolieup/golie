/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package models

import "github.com/rolieup/golie/internal/xml"

type JSONServiceRoot struct {
	Service Service `xml:"-" json:"service,"`
}

type Service struct {
	XMLName    xml.Name    `xml:"https://www.w3.org/2007/app service" json:"-"`
	Xmlns      string      `xml:"xmlns atom,attr" json:"-"`
	Lang       string      `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty" json:"-"`
	Workspaces []Workspace `xml:"workspace,omitempty" json:"workspace"`
}

type Workspace struct {
	Title       string       `xml:"atom title" json:"title"`
	Collections []Collection `xml:"collection,omitempty" json:"collection"`
}

type Collection struct {
	Href       string      `xml:"href,attr" json:"href"`
	Title      string      `xml:"atom title,omitempty" json:"title"`
	Link       *AtomLink   `xml:"atom link" json:"link"`
	Categories *Categories `xml:"categories,omitempty" json:"categories"`
}

type AtomLink struct {
	Link
}

type Categories struct {
	Fixed    string         `xml:"fixed,attr,omitempty" json:"fixed,omitempty"`
	Category []AtomCategory `xml:"atom category,omitempty" json:"category"`
}

type AtomCategory struct {
	Category
}
