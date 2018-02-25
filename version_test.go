package main

import (
	"os"
	"testing"
)

func TestExtractVersion(t *testing.T) {
	cases := []struct {
		input   string
		version string
	}{
		{"1.0.0", "1.0.0"},
		{"4.2.1\n Apple LLVM version 9.0.0", "9.0.0"},
		{"2", "2"},
		{"-", ""},
	}

	for _, tc := range cases {
		v, err := ExtractVersion(tc.input)
		if err != nil && tc.input != "-" {
			t.Fatalf("expected error for input: %s", tc.input)
		}

		if v != tc.version {
			t.Fatalf("input %s\nexpected %s\nactual: %s", tc.input, tc.version, v)
		}
	}
}

func TestMain(m *testing.M) {
	os.Args = []string{">1.0.0", "1.5.0"}
	os.Exit(m.Run())
}
