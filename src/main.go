package main

import (
	"fmt"
	"os"
	"time"

	"github.com/OlgaDnepr/goexporter/src/export"
	"github.com/OlgaDnepr/goexporter/src/export/entities"
	"github.com/OlgaDnepr/goexporter/src/export/formats"
)

func main() {
	var (
		filename = "userReport"
		user     = entities.User{
			Name:      "John",
			Age:       35,
			LastLogIn: time.Now(),
			WorkDay:   export.Friday,
		}
		dataHeaders1 = []string{"Name", "Age", "Age2"}
		dataHeaders2 = []string{"Name", "Last log in", "Age", "Work day"}
	)

	if err := exportToFile(filename, formats.NotationCSV, dataHeaders1, user); err != nil {
		fmt.Printf(" Error while exporting user %+v to %s format. Error: %s", user, formats.NotationCSV, err)
	}

	if err := exportToFile(filename, formats.NotationXLSX, dataHeaders2, user); err != nil {
		fmt.Printf(" Error while exporting user %+v to %s format. Error: %s", user, formats.NotationXLSX, err)
	}
}

func exportToFile(fileName, formatNotation string, dataHeaders []string, entities ...entities.Entity) error {
	fileNameFormat := fmt.Sprintf("%s.%s", fileName, formatNotation)
	file, err := os.Create(fileNameFormat)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error while closing file %s. Err: %s", fileNameFormat, err)
		}
	}()

	format, err := formats.New(formatNotation)
	if err != nil {
		return err
	}

	if err := format.SetUp(file, dataHeaders, entities...); err != nil {
		return err
	}
	return format.Write()
}
