package currency

import "testing"

func TestGetSymbol(t *testing.T) {
	testers := map[string]string{
		"CNY": "¥",
		"USD": "$",
	}

	for k, v := range testers {
		if v != GetSymbol(k) {
			t.Fatal("some error")
		}
	}
}
