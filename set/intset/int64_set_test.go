package intset

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewInt64Set()

	if set.Members() != nil {
		t.Fatalf("expected nil")
	}

	ms := []int64{1, 2, 3}
	set.Add(ms...)
	set.Add(ms...)
	if len(set.Members()) != 3 {
		t.Fatalf("expect members length is 3")
	}

	t.Logf("members: %v", set.Members())

	var u int64 = 5
	for _, v := range append(ms, u) {
		if v == u && set.HasMember(v) {
			t.Fatalf("%d is unexpected", v)
		}

		if v != u && !set.HasMember(v) {
			t.Fatalf("%d is expected", v)
		}
	}

	set.Del(ms[0], ms[2])
	if l := len(set.Members()); l != 1 {
		t.Fatalf("expected 2 but get: %d", l)
	} else if v := set.Members()[0]; v != ms[1] {
		t.Fatalf("expected %d but get: %d", ms[1], v)
	}

	t.Logf("members: %v", set.Members())

	set.Del(ms[1])
	set.Del(ms[1])
	set.Del(ms[1])
}

func TestInt64Set_Intersection(t *testing.T) {
	set1 := NewInt64Set(1, 2, -3, 4)
	set2 := NewInt64Set(3, 5, 6, 7, 8)

	v := set1.Intersection(set2)
	t.Logf("v: %v", v)
}
