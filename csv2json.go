package csv2json

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadCSV to read the content of CSV File
func ReadCSV(csvFile *os.File, additionalFields map[string]string, opt Options) ([]byte, error) {
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(content) < 1 {
		return nil, fmt.Errorf("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	if opt.LineWiseJSON {
		return lineWiseJSON(content, additionalFields, opt)
	}
	return intoJSONArray(content, additionalFields, opt)
}

// SaveFile creates a file by a given content and path
func SaveFile(content []byte, path string) error {
	return ioutil.WriteFile(path, content, os.FileMode(0644))
}

// ReadFile reads a file and creates a file handle
func ReadFile(path *string) (*os.File, error) {
	csvFile, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	return csvFile, nil
}
