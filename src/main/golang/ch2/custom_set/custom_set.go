package custom_set

import "iter"

type Set[T comparable] map[T]struct{}

func New[T comparable]() *Set[T] {
	return new(Set[T])
}

func (set *Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range *set {
			if !yield(elem) {
				return
			}
		}
	}
}

func (set *Set[T]) Clone() *Set[T] {
	clone := New[T]()
	for elem := range *set {
		clone.Add(elem)
	}
	return clone
}

func Collect[T comparable](it iter.Seq[T]) *Set[T] {
	newSet := New[T]()
	for t := range it {
		newSet.Add(t)
	}
	return newSet
}

func (set *Set[T]) Add(value T) (*Set[T], bool) {
	_, ok := (*set)[value]
	if ok {
		return set, false
	}
	(*set)[value] = struct{}{}
	return set, true
}

func (set *Set[T]) Remove(value T) (*Set[T], bool) {
	_, ok := (*set)[value]
	if !ok {
		return set, false
	}
	delete(*set, value)
	return set, true
}

func (set *Set[T]) Contains(value T) bool {
	_, ok := (*set)[value]
	return ok
}

func (set *Set[T]) Len() uint {
	return uint(len(*set))
}

func (set *Set[T]) Empty() bool {
	return len(*set) == 0
}

func (set *Set[T]) Clear() *Set[T] {
	clear(*set)
	return set
}

func (set *Set[T]) Union(other *Set[T]) *Set[T] {
	newSet := New[T]()
	for elem := range *set {
		newSet.Add(elem)
	}
	for elem := range *other {
		newSet.Add(elem)
	}
	return newSet
}

func (set *Set[T]) Intersection(other *Set[T]) *Set[T] {
	self := set
	if other.Len() < self.Len() {
		self, other = other, self
	}
	newSet := New[T]()
	for elem := range *self {
		if other.Contains(elem) {
			newSet.Add(elem)
		}
	}
	return newSet
}

func (set *Set[T]) Difference(other *Set[T]) *Set[T] {
	newSet := New[T]()
	for elem := range *set {
		if !other.Contains(elem) {
			newSet.Add(elem)
		}
	}
	return newSet
}

func (set *Set[T]) IsSubsetOf(other *Set[T]) bool {
	if set.Len() > other.Len() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *Set[T]) IsEqualTo(other *Set[T]) bool {
	if set.Len() != other.Len() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *Set[T]) IsProperSubsetOf(other *Set[T]) bool {
	if set.Len() >= other.Len() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *Set[T]) IsDisjointWith(other *Set[T]) bool {
	self := set
	if other.Len() < self.Len() {
		self, other = other, self
	}
	for elem := range *self {
		if other.Contains(elem) {
			return false
		}
	}
	return true
}
