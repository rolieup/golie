package models

import (
	"fmt"
)

const (
	Atom2005HttpsUri       = "https://www.w3.org/2005/Atom"
	Atom2005HttpUri        = "http://www.w3.org/2005/Atom"
	AtomPublishingHttpsUri = "https://www.w3.org/2007/app"
	AtomPublishingHttpUri  = "http://www.w3.org/2007/app"
	Roliens                = "urn:ietf:params:xml:ns:rolie-1.0"
)

type RolieRootElement interface {
	MarshalXMLRootPrepare()
}

// Prepare xml root (top-level) element for marshaling
func (s *Service) MarshalXMLRootPrepare() {
	if s.Xmlns == "" {
		s.Xmlns = AtomPublishingHttpsUri
	}
	if s.Atomns == "" {
		s.Atomns = Atom2005HttpsUri
	}
}

// Prepare xml root (top-level) element for marshaling
func (f *Feed) MarshalXMLRootPrepare() {
	if f.Xmlns == "" {
		f.Xmlns = Atom2005HttpsUri
	}
	if f.Roliens == "" {
		f.Roliens = Roliens
	}
	if f.Lang == "" {
		f.Lang = "en-US"
	}
}

// Prepare xml root (top-level) element for marshaling
func (e *Entry) MarshalXMLRootPrepare() {
	if e.Xmlns == "" {
		e.Xmlns = Atom2005HttpsUri
	}
	if e.Roliens == "" {
		e.Roliens = Roliens
	}
	if e.Lang == "" {
		e.Lang = "en-US"
	}
}

func AssertAtomNamespace(namespace string) error {
	switch namespace {
	case Atom2005HttpsUri:
		fallthrough
	case Atom2005HttpUri:
		return nil
	default:
		return fmt.Errorf("Unknown xml namespace '%s' expected %s", namespace, Atom2005HttpsUri)
	}
}

func AssertAtomPublishingNamespace(namespace string) error {
	switch namespace {
	case AtomPublishingHttpsUri:
		fallthrough
	case AtomPublishingHttpUri:
		return nil
	default:
		return fmt.Errorf("Unknown xml namespace '%s' expected %s", namespace, AtomPublishingHttpsUri)
	}
}
