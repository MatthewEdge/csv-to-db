package main

// PostgresTyper provides Postgres types to string input
type PostgresTyper struct {
	DefaultTyper
}

var typeWeight = map[string]int{
	"int":               1,
	"double":            2,
	"bool":              3,
	"character varying": 4,
}

// GetType to extract one of the supported types (double, int, bool, text) from the given input
func (t PostgresTyper) GetType(in string) string {
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

// MostWeight compares two types and returns the more likely of the two based on weight
func (t PostgresTyper) MostWeight(left, right string) string {
	if left == right {
		return left
	}

	if typeWeight[left] > typeWeight[right] {
		return left
	}

	return right
}

// GetLikelyType tries to extract the most likely type based on a set of input values
func (t PostgresTyper) GetLikelyType(in []string) string {

	// Exit early for single element lists
	if len(in) == 1 {
		return t.GetType(in[0])
	}

	var likelyType = t.GetType(in[0])

	for _, e := range in {
		resolved := t.GetType(e)
		likelyType = t.MostWeight(likelyType, resolved)
	}

	return likelyType
}
