package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	serverFlag := flag.Bool("server", false, "Start the app as a HTTP server?")
	pathFlag := flag.String("path", "", "Full path to the CSV file to parse")
	flag.Parse()

	if *serverFlag {
		server()
		os.Exit(0)
	}

	csv, err := FromFile(*pathFlag)

	if err != nil {
		fmt.Printf("Failed to convert CSV file: %s", err)
		os.Exit(1)
	}

	s := &Sqlizer{
		Typer:   PostgresTyper{},
		Headers: csv.Headers,
		Rows:    csv.Rows,
	}

	ddl := s.MakeDDL(csv.Name)
	inserts := s.MakeInserts(csv.Name)

	fmt.Printf("%s\n%s", ddl, strings.Join(inserts, "\n"))
	os.Exit(0)
}

func server() {
	http.HandleFunc("/load/postgres", func(w http.ResponseWriter, r *http.Request) {
		csv, err := FromHTTP(r)

		if err != nil {
			fmt.Fprint(w, "Failed to parse given CSV file")
			return
		}

		s := &Sqlizer{
			Typer:   PostgresTyper{},
			Headers: csv.Headers,
			Rows:    csv.Rows,
		}

		ddl := s.MakeDDL(csv.Name)
		inserts := strings.Join(s.MakeInserts(csv.Name), "\n")

		fmt.Fprintf(w, "%s\n%s", ddl, inserts)
	})

	http.HandleFunc("/load/file", func(w http.ResponseWriter, r *http.Request) {
		csv, err := FromFile("./test.csv")

		if err != nil {
			fmt.Fprint(w, "Failed to parse given CSV file")
			return
		}

		s := &Sqlizer{
			Typer:   PostgresTyper{},
			Headers: csv.Headers,
			Rows:    csv.Rows,
		}

		ddl := s.MakeDDL(csv.Name)
		inserts := strings.Join(s.MakeInserts(csv.Name), "\n")

		fmt.Fprintf(w, "%s\n%s", ddl, inserts)
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
