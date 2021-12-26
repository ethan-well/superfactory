package aregexp

import "testing"

func TestIsNumber(t *testing.T) {
	testers := map[string]bool{
		"1212":  true,
		"dddd":  false,
		"0":     true,
		"11d44": false,
		"":      false,
	}

	for k, v := range testers {
		if IsNumber(k) != v {
			t.Fatalf("expected: %v, k: %s", v, k)
		}
	}
}
