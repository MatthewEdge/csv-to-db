package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	serverFlag := flag.Bool("server", false, "Start the app as a HTTP server?")
	pathFlag := flag.String("path", "", "Full path to the CSV file to parse")
	dbHostFlag := flag.String("host", "", "Hostname of the database to load the CSV to")
	dbPortFlag := flag.String("port", "", "Port of the database to load the CSV to")
	dbUserFlag := flag.String("user", "", "User to authenticate as who can run the DDL and DML statements")
	dbPassFlag := flag.String("pass", "", "Password for the provided user")
	// dbNameFlag := flag.String("db", "", "Database to connect to")

	flag.Parse()

	// Currently only support postgres. Could extract for other options
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", *dbHostFlag, *dbPortFlag, *dbUserFlag, *dbPassFlag))
	exitOnError(err)
	defer db.Close()

	if *serverFlag {
		server(db)
		os.Exit(0)
	}

	csv, err := FromFile(*pathFlag)
	exitOnError(err)
	loadCsv(csv, db)

	os.Exit(0)
}

// loadCsv will attempt to load the given CsvFile to the given DB instance
func loadCsv(csv *CsvFile, db *sql.DB) error {
	s := &Sqlizer{
		Typer:   PostgresTyper{},
		Headers: csv.Headers,
		Rows:    csv.Rows,
	}

	ddl := s.MakeDDL(csv.Name)
	inserts := s.MakeInserts(csv.Name)

	_, err := db.Query(ddl)
	if err != nil {
		return err
	}

	fmt.Printf("Created table %s\n", csv.Name)

	for _, ins := range inserts {
		fmt.Println("Running statement: ", ins)
		_, err := db.Query(ins)
		if err != nil {
			return err
		}
	}

	return nil
}

// server spins up a HTTP API to upload CSV files
func server(db *sql.DB) {
	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		csv, err := FromHTTP(r)

		if err != nil {
			fmt.Fprint(w, "Failed to parse given CSV file")
			return
		}

		err = loadCsv(csv, db)
		if err != nil {
			fmt.Fprint(w, "Failed to load given CSV file to the database")
			return
		}
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
}
