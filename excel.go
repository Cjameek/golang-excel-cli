package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/xuri/excelize/v2"
)

type ExcelWorksheetData struct {
	Order string
}

func GetJsonResponse() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read stdin:", err)
		os.Exit(1)
	}

	var input ExcelWorksheetData
	if err := json.Unmarshal(inputBytes, &input); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid JSON:", err)
		os.Exit(1)
	}
}

func GenerateExcel(filename string) error {
	// Create a new Excel file
	f := excelize.NewFile()

	sheet := "Sheet1"

	// Set cell values
	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Score")
	f.SetCellValue(sheet, "A2", "Alice")
	f.SetCellValue(sheet, "B2", 95)
	f.SetCellValue(sheet, "A3", "Bob")
	f.SetCellValue(sheet, "B3", 88)

	f.MergeCell(sheet, "A3", "J3")

	// Save the file to the specified output path
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save excel file %w", err)
	}

	return nil
}
