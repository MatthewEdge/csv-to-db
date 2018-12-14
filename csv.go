package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

// CsvFile represents a parsed CSV file
type CsvFile struct {
	Name    string
	Headers []string
	Rows    [][]string
}

// Regexp to extract filename from a HTTP request
var fileNameMatcher = regexp.MustCompile(`^.*filename="(.*).csv"?`)

// FromFile will create a parsed CsvFile instance from the given file path
func FromFile(path string) (*CsvFile, error) {
	csvFile, err := os.Open(path)

	if err != nil {
		fmt.Printf("Failed to read CSV file at path: %s", path)
		return nil, err
	}

	parsed, err := read(csv.NewReader(bufio.NewReader(csvFile)))

	if err != nil {
		fmt.Printf("Failed to parse csv: %s", err)
		return nil, err
	}

	// Grab filename and then drop unnecessary headers and footers
	matched := fileNameMatcher.FindStringSubmatch(parsed[1][0])

	if len(matched) < 1 {
		fmt.Printf("Failed to extract filename from CSV file")
		return nil, errors.New("Failed to extract filename from csv file")
	}

	tableName := fileNameMatcher.FindStringSubmatch(parsed[1][0])[1]
	csv := parsed[3 : len(parsed)-1]
	headers := csv[0]
	rows := csv[1:]

	return &CsvFile{
		Name:    tableName,
		Headers: headers,
		Rows:    rows,
	}, nil
}

// FromHTTP will create a parsed CsvFile instance from the given http.Request
func FromHTTP(req *http.Request) (*CsvFile, error) {
	parsed, err := read(csv.NewReader(req.Body))

	if err != nil {
		return nil, err
	}

	// Grab filename and then drop unnecessary headers and footers
	tableName := fileNameMatcher.FindStringSubmatch(parsed[1][0])[1]
	csv := parsed[3 : len(parsed)-1]
	headers := csv[0]
	rows := csv[1:]

	return &CsvFile{
		Name:    tableName,
		Headers: headers,
		Rows:    rows,
	}, nil
}

// read will read CSV input raw from the given Reader
func read(reader *csv.Reader) ([][]string, error) {
	// parse POST body as csv
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to parse row")
			return nil, err
		}

		// add record to result set
		results = append(results, record)
	}

	return results, nil
}
