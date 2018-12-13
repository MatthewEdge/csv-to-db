package main

import (
	"regexp"
	"strings"
)

var floatMatcher = regexp.MustCompile(`^-?[0-9]{1,}\.[0-9]{1,}$`)
var intMatcher = regexp.MustCompile(`^(-?[1-9]+\d*)$|^0$`) // Leading 0 is not considered an int
var boolMatcher = regexp.MustCompile(`^true|false$`)

// Typer object to extract a valid type from the given value
type Typer interface {
	GetType(string) string
}

// DefaultTyper provides type methods to be used by implementation classes
type DefaultTyper struct{}

// IsFloat matches given string and checks if it can be boxed as a Float value
func (t DefaultTyper) IsFloat(in string) bool {
	return floatMatcher.MatchString(in)
}

// IsInt matches given string and checks if it can be boxed as an Int value
func (t DefaultTyper) IsInt(in string) bool {
	return intMatcher.MatchString(in)
}

// IsBool matches given string and checks if it can be boxed as a Boolean value
func (t DefaultTyper) IsBool(in string) bool {
	return boolMatcher.MatchString(strings.ToLower(in))
}

// IsText matches given string and checks if it can be boxed as a raw Text value
func (t DefaultTyper) IsText(in string) bool {
	return !t.IsBool(in) && !t.IsFloat(in) && !t.IsInt(in) // a bit naive but does work in DB contexts
}
