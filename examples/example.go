package main

import (
	"flag"
	"log"

	"github.com/cseeger-epages/csv2json"
)

func main() {
	// define csv path and output path
	csvPath := flag.String("c", "./csv/data.csv", "path of the file")
	outputPath := flag.String("o", "./json/data.json", "path of the output file")
	flag.Parse()

	// read and convert csv
	fileBytes, err := csv2json.ReadCSV(csvPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = csv2json.SaveFile(fileBytes, *outputPath)
	if err != nil {
		log.Fatal(err)
	}
}