/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/
package models

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	XMLName     xml.Name     `xml:"feed" json:"-"`
	Xmlns       string       `xml:"xmlns,attr" json:"-"`
	Lang        string       `xml:"xml:lang,attr" json:"-"`
	Roliens     string       `xml:"xmlns:rolie,attr" json:"-"`
	ID          string       `xml:"id" json:"id"`
	Title       string       `xml:"title" json:"title"`
	Description string       `xml:"description,omitempty" json:"description,omitempty"`
	Author      *Person      `xml:"author" json:"author,omitempty"`
	Contributor *Contributor `xml:"contributor" json:"contributor,omitempty"`
	Generator   *Generator   `xml:"generator,omitempty" json:"generator,omitempty"`
	Rights      string       `xml:"rights,omitempty" json:"rights,omitempty"`
	Link        []Link       `xml:"link" json:"link"`
	Category    *Category    `xml:"category" json:"category"`
	Updated     TimeStr      `xml:"updated" json:"updated"`
	Entry       []*Entry     `xml:"entry" json:"entry"`
}

type Category struct {
	Scheme string `xml:"scheme,attr" json:"scheme"`
	Term   string `xml:"term,attr" json:"term"`
}

type Generator struct {
	URI     string `xml:"uri,attr" json:"uri"`
	Version string `xml:"version,attr" json:"version"`
	Body    string `xml:",chardata" json:"-"`
}

type Entry struct {
	XMLName   xml.Name  `xml:"entry" json:"-"`
	Xmlns     string    `xml:"xmlns,attr" json:"-"`
	ID        string    `xml:"id" json:"id"`
	Title     string    `xml:"title" json:"title"`
	Link      []Link    `xml:"link" json:"link,omitempty"`
	Published TimeStr   `xml:"published" json:"published"`
	Updated   TimeStr   `xml:"updated" json:"updated"`
	Author    *Person   `xml:"author,omitempty" json:"author,omitempty"`
	Summary   *Text     `xml:"summary,omitempty" json:"summary,omitempty"`
	Content   *Text     `xml:"content,omitempty" json:"content,omitempty"`
	Format    *Format   `xml:"urn:ietf:params:xml:ns:rolie-1.0 format,omitempty" json:"format,omitempty"`
	Property  *Property `xml:"urn:ietf:params:xml:ns:rolie-1.0 property,omitempty" json:"property,omitempty"`
}

type Link struct {
	Rel      string `xml:"rel,attr,omitempty" json:"rel,omitempty"`
	Href     string `xml:"href,attr" json:"href"`
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty"`
	HrefLang string `xml:"hreflang,attr,omitempty" json:"-"`
	Title    string `xml:"title,attr,omitempty" json:"title,omitempty"`
	Length   uint   `xml:"length,attr,omitempty" json:"length,omitempty"`
	Body     string `xml:",chardata" json:"-"`
}

type Person struct {
	Name     string `xml:"name" json:"name"`
	URI      string `xml:"uri,omitempty" json:"uri,omitempty"`
	Email    string `xml:"email,omitempty" json:"email,omitempty"`
	Org      string `xml:"organization,omitempty" json:"organization,omitempty"`
	InnerXML string `xml:",innerxml" json:"-"`
}

type Contributor struct {
	Person
}

type Text struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Src  string `xml:"src,attr,omitempty" json:"src,omitempty"`
	Body string `xml:",chardata" json:"content,omitempty"`
}

type Format struct {
	Ns      string `xml:"ns,attr,omitempty" json:"schema,omitempty"`
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
}

type Property struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type TimeStr string

func Time(t time.Time) TimeStr {
	return TimeStr(t.Format(time.RFC3339))
}
