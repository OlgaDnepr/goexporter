package formats

import (
	"fmt"
	"strings"

	"github.com/OlgaDnepr/goexporter/src/export/entities"
	"github.com/pkg/errors"
)

// dataConverter converts entities' fields to a matrix of interface{}
// result contains dataHeaders as a first row
// entities are expected to have the same type
func dataConverter(dataHeaders []string, entities ...entities.Entity) ([][]interface{}, error) {
	if len(entities) == 0 {
		return nil, errors.New("no entity provided to export")
	}
	// convert []string to []interface{} and validate headers
	dataHeadersInterface := dataHeadersConvertAndValidate(dataHeaders, entities[0])
	data := make([][]interface{}, 0, len(entities)+1)
	data = append(data, dataHeadersInterface)
	for _, entity := range entities {
		row := make([]interface{}, 0, len(dataHeadersInterface))
		for _, header := range dataHeadersInterface {
			cell, _ := entity.FieldByHeader(fmt.Sprint(header))
			row = append(row, cell)
		}
		data = append(data, row)
	}
	return data, nil
}

// dataConverterAsString converts entities' fields to a matrix of string
// result contains dataHeaders as a first row
// entities are expected to have the same type
func dataConverterAsString(dataHeaders []string, entities ...entities.Entity) ([][]string, error) {
	dataInterface, err := dataConverter(dataHeaders, entities...)
	if err != nil {
		return nil, err
	}
	dataStr := make([][]string, 0, len(dataInterface))
	for _, row := range dataInterface {
		rowStr := make([]string, 0, len(row))
		for _, cell := range row {
			rowStr = append(rowStr, fmt.Sprint(cell))
		}
		dataStr = append(dataStr, rowStr)
	}
	return dataStr, nil
}

// dataHeadersConvertAndValidate converts dataHeaders []string to []interface{} and validate them
func dataHeadersConvertAndValidate(dataHeaders []string, entity entities.Entity) []interface{} {
	// convert []string to []interface{} and validate headers
	dataHeadersInterface := make([]interface{}, 0, len(dataHeaders))
	for _, header := range dataHeaders {
		if len(strings.TrimSpace(header)) == 0 {
			continue
		}
		// if the header is empty or invalid - just skip it
		if _, ok := entity.FieldByHeader(header); ok {
			dataHeadersInterface = append(dataHeadersInterface, header)
		}
	}
	return dataHeadersInterface
}
