package main

import "testing"

func TestGetType(t *testing.T) {
	typer := PostgresTyper{}

	var typeTests = []struct {
		in  string
		out string
	}{
		{"true", "bool"},
		{"false", "bool"},
		{"1", "int"},
		{"0", "int"},
		{"1234567890", "int"},
		{"0.0", "double"},
		{"123456789.123456789", "double"},
		{"94586709857.01", "double"},
		{"874859872698.1", "double"},
		{"yes", "character varying"},
		{"Welcome to ReAlLy FUN input", "character varying"},
		{" ", "character varying"},
		{"", "character varying"},
	}

	for _, tt := range typeTests {
		t.Run(tt.in, func(t *testing.T) {
			r := typer.GetType(tt.in)
			if r != tt.out {
				t.Errorf("got type %q for %q but expected %q", r, tt.in, tt.out)
			}
		})
	}
}

func TestGetLikelyType(t *testing.T) {
	typer := PostgresTyper{}

	var typeTests = []struct {
		in  []string
		out string
	}{
		{[]string{"true"}, "bool"},
		{[]string{"false", "true"}, "bool"},
		{[]string{"1"}, "int"},
		{[]string{"0", "1", "23456789"}, "int"},
		{[]string{"1", "0.0", "1.4567", "-1.6"}, "double"},
		{[]string{"0.0", "1", "welcome"}, "character varying"},
	}

	for _, tt := range typeTests {
		t.Run(tt.out, func(t *testing.T) {
			r := typer.GetLikelyType(tt.in)

			if r != tt.out {
				t.Errorf("got type %q for %q but expected %q", r, tt.in, tt.out)
			}
		})
	}
}
