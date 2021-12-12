package stringSet

import "sync"

type StringSet struct {
	values *sync.Map
}

func NewSet(values ...string) *StringSet {
	var set sync.Map
	for _, v := range values {
		set.Store(v, "")
	}

	return &StringSet{values: &set}
}

func (s *StringSet) Add(values ...string) {
	if s.values == nil {
		s.values = &sync.Map{}
	}
	for _, v := range values {
		s.values.Store(v, "")
	}
}

func (s *StringSet) Del(values ...string) {
	if s.values == nil {
		return
	}

	for _, v := range values {
		s.values.Delete(v)
	}
}

func (s *StringSet) HasMember(v string) bool {
	if s.values == nil {
		return false
	}

	_, ok := s.values.Load(v)
	return ok
}

func (s *StringSet) Members() []string {
	if s.values == nil {
		return nil
	}

	var values []string
	s.values.Range(func(key, value interface{}) bool {
		v, ok := key.(string)
		if ok {
			values = append(values, v)
		}
		return true
	})

	return values
}

func (s *StringSet) Clear() {
	s.values = nil
}
