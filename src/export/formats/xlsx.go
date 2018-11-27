package formats

import (
	"io"

	"github.com/OlgaDnepr/goexporter/src/export/entities"
	"github.com/tealeg/xlsx"
)

// FormatXLSX is the implementation of Format interface to represent execution results in XLSX format
type FormatXLSX struct {
	w    io.Writer
	data [][]interface{}
}

const sheetName = `Sheet1`

// SetUp fills in all fields of FormatXLSX by provided parameters
func (f *FormatXLSX) SetUp(w io.Writer, dataHeaders []string, entities ...entities.Entity) (err error) {
	f.data, err = dataConverter(dataHeaders, entities...)
	if err != nil {
		return err
	}
	f.w = w
	return nil
}

// Write writes entities in xlsx format
func (f *FormatXLSX) Write() error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}
	for _, row := range f.data {
		sheet.AddRow().WriteSlice(&row, -1)
	}
	return file.Write(f.w)
}
