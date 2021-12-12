package stringSet

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet()

	if set.Members() != nil {
		t.Fatalf("expected nil")
	}

	ms := []string{"v1", "v2", "v3"}
	set.Add(ms...)
	set.Add(ms...)
	if len(set.Members()) != 3 {
		t.Fatalf("expect members length is 3")
	}

	t.Logf("members: %v", set.Members())

	u := "v5"
	for _, v := range append(ms, u) {
		if v == u && set.HasMember(v) {
			t.Fatalf("%s is unexpected", v)
		}

		if v != u && !set.HasMember(v) {
			t.Fatalf("%s is expected", v)
		}
	}

	set.Del(ms[0], ms[2])
	if l := len(set.Members()); l != 1 {
		t.Fatalf("expected 2 but get: %d", l)
	} else if v := set.Members()[0]; v != ms[1] {
		t.Fatalf("expected %s but get: %s", ms[1], v)
	}

	t.Logf("members: %v", set.Members())

	set.Del(ms[1])
	set.Del(ms[1])
	set.Del(ms[1])
}
