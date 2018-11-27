package formats

import (
	"fmt"
	"io"

	"github.com/OlgaDnepr/goexporter/src/export/entities"
)

// Format is a generic interface to define different formats to export
type Format interface {
	SetUp(w io.Writer, dataHeaders []string, entities ...entities.Entity) error
	Write() error
}

// These constants define possible formats for export to
const (
	NotationCSV  = "csv"
	NotationXLSX = "xlsx"
)

// New creates a new format based on the provided format notation
func New(formatNotation string) (format Format, err error) {
	switch formatNotation {
	case NotationXLSX:
		format = &FormatXLSX{}
	case NotationCSV:
		format = &FormatCSV{}
	default:
		return format, fmt.Errorf("incorrect format: %q", formatNotation)
	}
	return
}
