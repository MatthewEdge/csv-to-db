package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
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

	http.ListenAndServe("0.0.0.0:8080", nil)
}
