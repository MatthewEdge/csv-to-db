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

	ddl := `CREATE TABLE person (
	first_name character varying,
	last_name character varying,
	age int,
	latitude double,
	longitude double,
	owns_dog bool,
	license_number character varying
);`

	if s.MakeDDL("person") != ddl {
		t.Errorf("MakeDDL did not produce the correct DDL statement. Got:\n%s", ddl)
	}
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
