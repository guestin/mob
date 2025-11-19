package mob

import mapset "github.com/deckarep/golang-set"

func NewSet() mapset.Set {
	return mapset.NewThreadUnsafeSet()
}

func NewSetFrom[T any](vs ...T) mapset.Set {
	return NewSetFromSlice(vs)
}

func NewSetFromSlice[T any](vs []T) mapset.Set {
	set := mapset.NewThreadUnsafeSet()
	for _, item := range vs {
		set.Add(item)
	}
	return set
}

func NewConcurrentSet() mapset.Set {
	return mapset.NewSet()
}

func NewConcurrentSetFrom[T any](vs ...T) mapset.Set {
	return NewConcurrentSetFromSlice(vs)
}

func NewConcurrentSetFromSlice[T any](vs ...T) mapset.Set {
	set := mapset.NewSet()
	for _, item := range vs {
		set.Add(item)
	}
	return set
}

func SetAdd(dst mapset.Set, vs ...interface{}) mapset.Set {
	return SetFromSlice(dst, vs)
}

func SetFromSlice(dst mapset.Set, vs []interface{}) mapset.Set {
	for _, it := range vs {
		dst.Add(it)
	}
	return dst
}
