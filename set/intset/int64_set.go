package intset

import "sync"

type Int64Set struct {
	values *sync.Map
}

func NewInt64Set(vs ...int64) *Int64Set {
	var set sync.Map
	for _, v := range vs {
		set.Store(v, "")
	}

	return &Int64Set{values: &set}
}

func (s *Int64Set) Add(vs ...int64) {
	if s.values == nil {
		s.values = &sync.Map{}
	}
	for _, v := range vs {
		s.values.Store(v, "")
	}
}

func (s *Int64Set) Del(vs ...int64) {
	if s.values == nil {
		return
	}

	for _, v := range vs {
		s.values.Delete(v)
	}
}

func (s *Int64Set) HasMember(v int64) bool {
	if s.values == nil {
		return false
	}

	_, ok := s.values.Load(v)
	return ok
}

func (s *Int64Set) Members() []int64 {
	if s.values == nil {
		return nil
	}

	var values []int64
	s.values.Range(func(key, value interface{}) bool {
		v, ok := key.(int64)
		if ok {
			values = append(values, v)
		}
		return true
	})

	return values
}

func (s *Int64Set) Clear() {
	s.values = nil
}

func (s *Int64Set) Intersection(set *Int64Set) []int64 {
	var values []int64
	s.values.Range(func(key, value interface{}) bool {
		v, _ := key.(int64)

		if set.HasMember(v) {
			values = append(values, v)
		}

		return true
	})

	return values
}
