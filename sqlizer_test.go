package main

import (
	"fmt"
	"testing"
)

func TestSqlizer_MakeDDL(t *testing.T) {
	s := &Sqlizer{
		Typer:   PostgresTyper{},
		Headers: []string{"first_name", "last_name", "age", "latitude", "longitude", "owns_dog", "license_number"},
		Rows: [][]string{
			{"steven", "smith", "40", "33.749", "-84.3880", "true", "1355322"},
			{"jane", "austen", "33", "33.751", "-84.234", "false", "ZA434324"},
		},
	}

	generated := s.MakeDDL("person")
	fmt.Println(generated)
}

func TestSqlizer_MakeInserts(t *testing.T) {
	s := &Sqlizer{
		Typer:   PostgresTyper{},
		Headers: []string{"first_name", "last_name", "age", "latitude", "longitude", "owns_dog", "license_number"},
		Rows: [][]string{
			{"steven", "smith", "40", "33.749", "-84.3880", "true", "1355322"},
			{"jane", "austen", "33", "33.751", "-84.234", "false", "ZA434324"},
		},
	}

	for _, i := range s.MakeInserts("person") {
		fmt.Println(i)
	}
}
