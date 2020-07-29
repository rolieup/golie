package models

import (
	"fmt"
)

const (
	Atom2005HttpsUri       = "https://www.w3.org/2005/Atom"
	Atom2005HttpUri        = "http://www.w3.org/2005/Atom"
	AtomPublishingHttpsUri = "https://www.w3.org/2007/app"
	AtomPublishingHttpUri  = "http://www.w3.org/2007/app"
)

// Prepare xml root (top-level) element for marshaling
func (s *Service) MarshalXMLRootPrepare() {
	s.Xmlns = AtomPublishingHttpsUri
}

// Prepare xml root (top-level) element for marshaling
func (f *Feed) MarshalXMLRootPrepare() {
	f.Xmlns = Atom2005HttpsUri
}

// Prepare xml root (top-level) element for marshaling
func (e *Entry) MarshalXMLRootPrepare() {
	e.Xmlns = Atom2005HttpsUri
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
