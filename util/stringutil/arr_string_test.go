package stringutil

import "testing"

func TestArrStringToInt64Arr(t *testing.T) {
	testers := map[string]map[bool][]int64{
		"[1,2,3,4]": {true: []int64{1, 2, 3, 4}},
		"[1]":       {true: []int64{1}},
		"[]":        {true: []int64{}},
		"":          {true: []int64{}},
		"1,2,3,4":   {false: nil},
		"[1,2,3":    {false: nil},
		"1,2,3]":    {false: nil},
	}

	for s, e := range testers {

		for r, arr := range e {
			result, err := ArrStringToInt64Arr(s)
			if (err == nil) != r {
				t.Fatal(err, "s: ", s)
			}
			if len(result) != len(arr) {
				t.Fatalf("expeced: %v, get: %v", arr, result)
			}

			for i, v := range arr {
				if result[i] != v {
					t.Fatalf("index: %d, expected: %d, get: %d", i, v, result[i])
				}
			}
		}
	}
}
