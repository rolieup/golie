package rolie

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocomply/scap/pkg/scap/constants"
	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/rolieup/golie/pkg/models"
	"github.com/rolieup/golie/pkg/rolie_source"
)

type Builder struct {
	ID            string
	Title         string
	RootURI       string
	DirectoryPath string
}

func (b *Builder) New() error {
	feed, err := b.feedForDirectory()
	if err != nil {
		return err
	}
	doc := rolie_source.Document{
		Feed: feed,
	}
	return doc.Write(filepath.Join(b.DirectoryPath, "feed"))
}

func (b *Builder) feedForDirectory() (*models.Feed, error) {
	scapFiles, err := traverseScapFiles(b.DirectoryPath)
	if err != nil {
		return nil, err
	}
	feed := models.Feed{
		ID:      b.ID,
		Title:   b.Title,
		Updated: models.Time(time.Now()),
	}
	for f := range scapFiles {
		entry, err := f.RolieEntry(b.RootURI)
		if err != nil {
			return nil, err
		}
		feed.Entry = append(feed.Entry, entry)
	}
	return &feed, nil
}

type scapFile struct {
	Path string
	*scap_document.Document
	Size         int64
	ModifiedTime time.Time
}

func (scap *scapFile) RolieEntry(baseUri string) (*models.Entry, error) {
	var entry models.Entry
	var err error

	entry.ID = scap.Document.Type.ShortName() + ":" + scap.Path
	entry.Title, err = scap.Title()
	if err != nil {
		return nil, err
	}
	entry.Link = []models.Link{
		models.Link{
			Href:   scap.Link(baseUri),
			Length: uint64(scap.Size),
		},
	}
	entry.Updated = models.Time(scap.ModifiedTime)
	entry.Published = models.Time(time.Now())
	entry.Content = &models.Text{
		Type: "applicatiton/xml",
		Src:  scap.Link(baseUri),
	}
	entry.Format = &models.Format{
		Ns:      scap.Document.Xmlns(),
		Version: scap.Document.ScapVersion(),
	}
	return &entry, nil
}

func (scap *scapFile) Link(baseUri string) string {
	baseUri = strings.TrimSuffix(baseUri, "/")
	path := strings.TrimPrefix(scap.Path, "/")
	return baseUri + "/" + path
}

func (scap *scapFile) Title() (string, error) {
	switch scap.Document.Type {
	case constants.DocumentTypeXccdfBenchmark:
		if len(scap.Document.Benchmark.Title) > 0 {
			return scap.Document.Benchmark.Title[0].Text, nil
		}
	case constants.DocumentTypeCpeDict:
		return "CPE Dictionary", nil
	case constants.DocumentTypeOcil:
		return "OCIL Questionaire", nil
	case constants.DocumentTypeOvalDefinitions:
		classes := scap.Document.OvalDefinitions.DefinitionClasses()
		switch len(classes) {
		case 0:
			return "Empty OVAL Definitions", nil
		case 1:
			return "OVAL " + strings.Title(string(classes[0])) + " Definitions", nil
		default:
			return "OVAL Definitions", nil
		}
	case constants.DocumentTypeSourceDataStream:
		datastreams := scap.Document.DataStreamCollection.DataStream
		switch len(datastreams) {
		case 0:
			return "Empty SCAP DataStream Collection", nil
		case 1:
			if datastreams[0].Checklists != nil {
				checklists := datastreams[0].Checklists.ComponentRef
				switch len(checklists) {
				case 0:
					return "SCAP DataStream without checklists", nil
				case 1:
					component := scap.Document.DataStreamCollection.GetComponentByRef(&checklists[0])
					if component != nil && component.DocumentType() == constants.DocumentTypeXccdfBenchmark {
						if len(component.Benchmark.Title) > 0 {
							return "SCAP DataStream for " + component.Benchmark.Title[0].Text, nil
						}
					}
				default:
					return "SCAP DataStream with multiple checklists", nil
				}
			}
		default:
			return "SCAP DataStream Collection", nil
		}

	}
	return "", fmt.Errorf("Not implemented: could not determine sensible atom:title for document '%s'", scap.Path)
}

func traverseScapFiles(directoryPath string) (<-chan scapFile, error) {
	out := make(chan scapFile)

	var err error
	go func() {
		err = filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("Internal error: could not walk the filesystem: %v", err)
			}
			if info.IsDir() || !strings.HasSuffix(path, ".xml") {
				return nil
			}

			doc, err := scap_document.ReadDocumentFromFile(path)
			if err != nil {
				return nil
			}
			out <- scapFile{
				Path:         strings.TrimPrefix(path, directoryPath),
				Document:     doc,
				Size:         info.Size(),
				ModifiedTime: info.ModTime(),
			}
			return nil
		})
		close(out)
	}()
	return out, err
}
