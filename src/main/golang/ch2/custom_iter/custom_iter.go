// Package custom_iter is my Golang implementation of Iterator in Rust
// nightly-only experimental API was excluded from this implementation.
// Methods that turn iterator into collection(s) such as collect & partition were limited to slice rather than generic "collection".
// In my opinion in Golang, these methods need to be implemented on per collection type basis.
// In Rust there is FromIterator trait which defines how conversion from Iterator to each Collection works.
// However, in Golang such unified interface defining iterator -> collection conversion doesn't exist.
package custom_iter

import (
	"iter"
	"slices"
)

func Count[T any](it iter.Seq[T]) uint {
	count := uint(0)
	for range it {
		count++
	}
	return count
}

func Last[T any](it iter.Seq[T]) (T, bool) {
	ok := false
	var elem T
	for t := range it {
		ok = true
		elem = t
	}
	return elem, ok
}

func Nth[T any](it iter.Seq[T], n uint) (T, bool) {
	var elem T
	ok := false
	count := uint(0)
	for t := range it {
		if count == n {
			ok = true
			elem = t
		}
		count++
	}
	return elem, ok
}

func StepBy[T any](it iter.Seq[T], step uint) iter.Seq[T] {
	if step == 0 {
		panic("StepBy step cannot be zero")
	}
	return func(yield func(t T) bool) {
		index := uint(0)
		for t := range it {
			if index%step == 0 {
				if !yield(t) {
					return
				}
			}
			index++
		}
	}
}

func Chain[T any](it, other iter.Seq[T]) iter.Seq[T] {
	return func(yield func(t T) bool) {
		for t := range it {
			if !yield(t) {
				return
			}
		}
		for o := range other {
			if !yield(o) {
				return
			}
		}
	}
}

func Zip[T, O any](it iter.Seq[T], other iter.Seq[O]) iter.Seq2[T, O] {
	return func(yield func(t T, o O) bool) {
		next, stop := iter.Pull(other)
		defer stop()
		for t := range it {
			o, ok := next()
			if !ok {
				return
			}
			if !yield(t, o) {
				return
			}
		}
	}
}

func ForEach[T any](it iter.Seq[T], doFn func(T)) {
	for t := range it {
		doFn(t)
	}
}

func All[T any](it iter.Seq[T], pred func(T) bool) bool {
	for t := range it {
		if !pred(t) {
			return false
		}
	}
	return true
}

func Any[T any](it iter.Seq[T], pred func(T) bool) bool {
	for t := range it {
		if pred(t) {
			return true
		}
	}
	return false
}

func ByRef[T any](it iter.Seq[T]) iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for t := range it {
			if !yield(&t) {
				return
			}
		}
	}
}

func CollectIntoSlice[T any](it iter.Seq[T]) []T {
	return slices.Collect[T](it)
}

func PartitionIntoSlices[T any](it iter.Seq[T], pred func(T) bool) ([]T, []T) {
	var satisfies, rest []T
	for t := range it {
		if pred(t) {
			satisfies = append(satisfies, t)
		} else {
			rest = append(rest, t)
		}
	}
	return satisfies, rest
}

func TryFold[T, R any, E error](it iter.Seq[T], init R, foldFn func(R, T) (R, E)) (R, E) {
	acc := init
	var err E
	for t := range it {
		acc, err = foldFn(acc, t)
		if err != nil {
			return acc, err
		}
	}
	return acc, err
}

func TryForEach[T any, E error](it iter.Seq[T], doFn func(T) E) E {
	var err E
	for t := range it {
		err = doFn(t)
		if err != nil {
			return err
		}
	}
	return err
}

func Fold[T, R any](it iter.Seq[T], init R, foldFn func(R, T) R) R {
	acc := init
	for t := range it {
		acc = foldFn(acc, t)
	}
	return acc
}

func Reduce[T any](it iter.Seq[T], reduceFn func(T, T) T) (T, bool) {
	var acc T
	ok := false
	for t := range it {
		if !ok {
			acc = t
			ok = true
		} else {
			acc = reduceFn(acc, t)
		}
	}
	return acc, ok
}

func Eq[T comparable](it, other iter.Seq[T]) bool {
	next, stop := iter.Pull(other)
	defer stop()
	var o T
	var ok bool
	for t := range it {
		o, ok = next()
		if !ok {
			return false
		}
		if t != o {
			return false
		}
	}
	return true
}

func EqBy[T any](it, other iter.Seq[T], equalFn func(T, T) bool) bool {
	next, stop := iter.Pull(other)
	defer stop()
	var o T
	var ok bool
	for t := range it {
		o, ok = next()
		if !ok {
			return false
		}
		if !equalFn(o, t) {
			return false
		}
	}
	return true
}

func Filter[T any](it iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range it {
			if pred(t) {
				if !yield(t) {
					return
				}
			}
		}
	}
}

func Map[T, R any](it iter.Seq[T], mapFn func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		var r R
		for t := range it {
			r = mapFn(t)
			if !yield(r) {
				return
			}
		}
	}
}

func FilterMap[T, R any](it iter.Seq[T], filterMapFn func(T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		var r R
		var ok bool
		for t := range it {
			r, ok = filterMapFn(t)
			if ok {
				if !yield(r) {
					return
				}
			}
		}
	}
}

func Enumerate[T any](it iter.Seq[T]) iter.Seq2[uint, T] {
	return func(yield func(uint, T) bool) {
		var index uint
		for t := range it {
			if !yield(index, t) {
				return
			}
			index++
		}
	}
}

type PeekableSeq[T any] iter.Seq2[Option[T], Option[T]]

type Option[T any] struct {
	ok    bool
	value T
}

func (peekableSeq PeekableSeq[T]) Peek() (T, bool) {
	var value T
	var ok bool
	peekableSeq(func(first, second Option[T]) bool {
		value, ok = second.value, second.ok
		return false
	})
	return value, ok
}

func Peekable[T any](it iter.Seq[T]) PeekableSeq[T] {
	return func(yield func(Option[T], Option[T]) bool) {
		var prev T
		var notFirst bool
		for t := range it {
			if !notFirst {
				notFirst = true
				prev = t
				continue
			}
			if !yield(Option[T]{ok: true, value: prev}, Option[T]{ok: true, value: t}) {
				prev = t
				return
			}
			prev = t
		}
		if !yield(Option[T]{ok: true, value: prev}, Option[T]{ok: false}) {
			return
		}
	}
}

func SkipWhile[T any](it iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var stopSkip bool
		for t := range it {
			if !pred(t) {
				stopSkip = true
			}
			if !stopSkip {
				continue
			}
			if !yield(t) {
				return
			}
		}
	}
}

func TakeWhile[T any](it iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var stopTake bool
		for t := range it {
			if !pred(t) {
				stopTake = true
			}
			if stopTake {
				break
			}
			if !yield(t) {
				return
			}
		}
	}
}

func MapWhile[T, R any](it iter.Seq[T], mapWhileFn func(T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		var stopMap bool
		var r R
		var ok bool
		for t := range it {
			r, ok = mapWhileFn(t)
			if !ok {
				stopMap = true
			}
			if stopMap {
				break
			}
			if !yield(r) {
				return
			}
		}
	}
}

func Skip[T any](it iter.Seq[T], n uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		index := uint(0)
		for t := range it {
			if index >= n {
				if !yield(t) {
					return
				}
			}
			index++
		}
	}
}

func Take[T any](it iter.Seq[T], n uint) iter.Seq[T] {
	return func(yield func(T) bool) {
		index := uint(0)
		for t := range it {
			if index >= n {
				break
			}
			if !yield(t) {
				return
			}
			index++
		}
	}
}

func Scan[T, R any](it iter.Seq[T], init R, scanFn func(R, T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		acc := init
		var ok bool
		var r R
		for t := range it {
			r, ok = scanFn(acc, t)
			if !ok {
				break
			}
			acc = r
			if !yield(r) {
				return
			}
		}
	}
}

func FlatMap[T, R any](it iter.Seq[T], mapFn func(T) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		var innerIter iter.Seq[R]
		for t := range it {
			innerIter = mapFn(t)
			for r := range innerIter {
				if !yield(r) {
					return
				}
			}
		}
	}
}

func Flatten[T any](it iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for innerIter := range it {
			for t := range innerIter {
				if !yield(t) {
					return
				}
			}
		}
	}
}

func Fuse[T any](it iter.Seq[Option[T]]) iter.Seq[Option[T]] {
	return func(yield func(Option[T]) bool) {
		for t := range it {
			if !t.ok {
				break
			}
			if !yield(t) {
				return
			}
		}
	}
}

func Inspect[T any](it iter.Seq[T], doFn func(T)) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range it {
			doFn(t)
			if !yield(t) {
				return
			}
		}
	}
}

func Find[T any](it iter.Seq[T], pred func(T) bool) (T, bool) {
	for t := range it {
		if pred(t) {
			return t, true
		}
	}
	var zero T
	return zero, false
}

func FindMap[T, R any](it iter.Seq[T], findFn func(T) (R, bool)) (R, bool) {
	var r R
	var ok bool
	for t := range it {
		r, ok = findFn(t)
		if ok {
			return r, true
		}
	}
	var zero R
	return zero, false
}

//TODO: Position, RPosition, Max, Min, MaxByKey, MaxBy, MinByKey, MinBy, Rev, Unzip, Copied, Cloned, Cycle, Sum, Product, Cmp, PartialCmp, Ne, Lt, Le, Gt, Ge, IsSorted, IsSortedBy, IsSortedByKey
