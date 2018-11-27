package entities

import (
	"reflect"

	"github.com/OlgaDnepr/goexporter/src/export"
)

// Entity is a generic interface which must be satisfied by the entity to export
// Entity will be presented as a row in a result file
type Entity interface {
	FieldByHeader(header string) (interface{}, bool)
}

func fieldByHeaderEntity(entity Entity, header string) (interface{}, bool) {
	value := reflect.ValueOf(entity)
	for i := 0; i < value.NumField(); i++ {
		if header != value.Type().Field(i).Tag.Get("export-header") {
			continue
		}
		field := value.Field(i).Interface()
		// Special conversion for custom types like enums
		if ex, ok := field.(export.Exporter); ok {
			return ex.Export(), true
		}
		return field, true
	}
	return "", false
}
