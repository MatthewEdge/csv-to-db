package main

import "testing"

func TestIsFloat(t *testing.T) {
	m := DefaultTyper{}

	validFloats := []string{"1.0", "-1.0", "123.456", "-123.13456", "0.847562987645826478", "123456789098765432.2345678909876543212345678987654323456789"}

	for _, i := range validFloats {
		if !m.IsFloat(i) {
			t.Error(i + " should be flagged as a float")
		}
	}

	invalidFloats := []string{"abc", "1334245", "true", "", " "}
	for _, ni := range invalidFloats {
		if m.IsFloat(ni) {
			t.Error(ni + " should NOT be flagged as a float")
		}
	}
}

func TestIsInt(t *testing.T) {
	m := DefaultTyper{}

	validInts := []string{"1", "-1", "1232435465", "-1234567", "783587", "87654323456789876543212345"}

	for _, i := range validInts {
		if !m.IsInt(i) {
			t.Error(i + " should be flagged as an int")
		}
	}

	invalidInts := []string{"0123", "1a", "abc", "1334245.1213", "", " "}
	for _, ni := range invalidInts {
		if m.IsInt(ni) {
			t.Error(ni + " should NOT be flagged as an int")
		}
	}
}

func TestIsBool(t *testing.T) {
	m := DefaultTyper{}

	validBools := []string{"true", "false", "True", "False", "TRUE", "FALSE", "tRuE", "fAlSe"}

	for _, b := range validBools {
		if !m.IsBool(b) {
			t.Error(b + " should be flagged as a bool")
		}
	}

	invalidBools := []string{"1", "0", "abc", "1334245.1213", "", " "}
	for _, nb := range invalidBools {
		if m.IsBool(nb) {
			t.Error(nb + " should NOT be flagged as a bool")
		}
	}
}
