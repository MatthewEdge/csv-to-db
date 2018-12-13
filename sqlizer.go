package main

import (
	"fmt"
	"strings"
)

// Sqlizer converts parsed CSV input into SQL statements
type Sqlizer struct {
	Typer       Typer
	Headers     []string
	HeaderTypes []string
	Rows        [][]string
}

// MakeDDL produces the `CREATE TABLE` statement
func (s *Sqlizer) MakeDDL(tableName string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))

	transposed := transpose(s.Rows)
	for i, h := range s.Headers {
		column := transposed[i][0:]
		likelyType := s.Typer.GetLikelyType(column)

		sb.WriteString(fmt.Sprintf("\t%s %s", h, likelyType))

		// Commas for all but the last element
		if i < len(s.Headers)-1 {
			sb.WriteString(",")
		}

		sb.WriteString("\n")
	}

	sb.WriteString(");")

	return sb.String()
}

// MakeInserts produces the `INSERT` statements for each row given
func (s *Sqlizer) MakeInserts(tableName string) []string {
	inserts := make([]string, len(s.Rows))
	headers := strings.Join(s.Headers, ",")

	for i, r := range s.Rows {
		inserts[i] = fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s');", tableName, headers, strings.Join(r, "','"))
	}

	return inserts
}

// Transpose a slice of [a x b] to a slice of [b x a]
func transpose(slice [][]string) [][]string {
	aLen := len(slice[0])
	bLen := len(slice)
	transposed := make([][]string, aLen)

	for i := range transposed {
		transposed[i] = make([]string, bLen)
	}

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			transposed[i][j] = slice[j][i]
		}
	}

	return transposed
}
