package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
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

// FromHTTP will create a parsed CsvFile instance from the given http.Request
func FromHTTP(req *http.Request) (*CsvFile, error) {

	parsed, err := readHTTP(req)

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

// ReadHTTP will read CSV file input from a HTTP request
func readHTTP(req *http.Request) ([][]string, error) {
	// parse POST body as csv
	reader := csv.NewReader(req.Body)
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
