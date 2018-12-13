package main

// PostgresTyper provides Postgres types to string input
type PostgresTyper struct {
	DefaultTyper
}

// GetType to extract one of the supported types (double, int, bool, text) from the given input
func (t *PostgresTyper) GetType(in string) string {
	if t.IsFloat(in) {
		return "double"
	} else if t.IsInt(in) {
		return "int"
	} else if t.IsBool(in) {
		return "bool"
	} else {
		return "character varying"
	}
}
