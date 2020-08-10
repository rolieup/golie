package rolie

import (
	"fmt"
	"github.com/rolieup/golie/pkg/rolie_source"
)

func Info(uri string) error {
	document, err := rolie_source.ReadDocumentFromURI(uri)
	if err != nil {
		return fmt.Errorf("Failed to parse rolie document %s", err)
	}
	if document.Feed != nil {
		fmt.Println("Document Type: ROLIE Feed")
		feed := document.Feed
		if feed.Title != "" {
			fmt.Printf("Title: %s\n", feed.Title)
		}
		if feed.Updated != "" {
			fmt.Printf("Updated: %s\n", feed.Updated)
		}

		entries := "entries"
		if len(feed.Entry) == 1 {
			entries = "entry"
		}
		fmt.Printf("Contains %d %s.\n", len(feed.Entry), entries)
	} else if document.Service != nil {
		fmt.Println("Document Type: ROLIE Service")
		service := document.Service
		if len(service.Workspaces) == 0 {
			fmt.Println("Contains 0 workspaces.")
		} else {
			fmt.Println("Workspaces:")
			for _, w := range service.Workspaces {
				fmt.Printf("\t- Title: %s\n", w.Title)
				for _, c := range w.Collections {
					fmt.Printf("\t\t- Title: %s\n", c.Title)
					fmt.Printf("\t\t- Href: %s\n", c.Href)
				}
			}

		}
		fmt.Println("Workspaces")
	} else if document.Entry != nil {
		fmt.Println("Document Type: ROLIE Entry")
		entry := document.Entry
		if entry.Title != "" {
			fmt.Printf("Title: %s\n", entry.Title)
		}
		if entry.Published != "" {
			fmt.Printf("Published: %s\n", entry.Published)
		}
	} else {
		return fmt.Errorf("Could not recognize ROLIE resource")
	}
	return nil
}
