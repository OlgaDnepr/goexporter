package formats

import (
	"encoding/csv"

	"io"

	"github.com/OlgaDnepr/goexporter/src/export/entities"
)

// FormatCSV is the implementation of Format interface to represent execution results in CSV format
type FormatCSV struct {
	w    io.Writer
	data [][]string
}

// SetUp fills in all fields of FormatCSV by provided parameters
func (f *FormatCSV) SetUp(w io.Writer, dataHeaders []string, entities ...entities.Entity) (err error) {
	f.data, err = dataConverterAsString(dataHeaders, entities...)
	if err != nil {
		return err
	}
	f.w = w
	return nil
}

// Write writes entities in csv format
func (f *FormatCSV) Write() error {
	file := csv.NewWriter(f.w)
	return file.WriteAll(f.data)
}
