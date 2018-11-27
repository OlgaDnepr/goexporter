package entities

import (
	"time"

	"github.com/OlgaDnepr/goexporter/src/export"
)

// User contains user info for report
type User struct {
	Name      string            `export-header:"Name"`
	LastLogIn time.Time         `export-header:"Last log in"`
	WorkDay   export.DaysOfWeek `export-header:"Work day"`
	Age       int               `export-header:"Age"`
}

// FieldByHeader is an implementation of the Entity interface
func (u User) FieldByHeader(header string) (interface{}, bool) {
	return fieldByHeaderEntity(u, header)
}
