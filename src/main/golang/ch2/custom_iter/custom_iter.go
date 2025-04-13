// Package custom_iter is my Golang implementation of Iterator in Rust
// nightly-only experimental API (as of Rust 1.86.0) were excluded from this implementation.
// Methods that turn iterator into collection(s) such as collect & partition were limited to slice rather than generic "collection".
// In my opinion in Golang, these methods need to be implemented on per collection type basis.
// In Rust there is FromIterator trait which defines how conversion from Iterator to each Collection works.
// However, in Golang such unified interface defining iterator -> collection conversion doesn't exist.
package custom_iter

import (
	"cmp"
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
	var err error
	for t := range it {
		acc, err = foldFn(acc, t)
		if err != nil {
			return acc, err.(E)
		}
	}
	return acc, err.(E)
}

func TryForEach[T any, E error](it iter.Seq[T], doFn func(T) E) E {
	var err error
	for t := range it {
		err = doFn(t)
		if err != nil {
			return err.(E)
		}
	}
	return err.(E)
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

func Position[T any](it iter.Seq[T], pred func(T) bool) (uint, bool) {
	index := uint(0)
	for t := range it {
		if pred(t) {
			return index, true
		}
		index++
	}
	return 0, false
}

func RPosition[T any](it iter.Seq[T], pred func(T) bool) (uint, bool) {
	found := false
	index := uint(0)
	var at uint
	for t := range it {
		if pred(t) {
			found = true
			at = index
		}
		index++
	}
	if found {
		return at, true
	} else {
		return 0, false
	}
}

func Max[T cmp.Ordered](it iter.Seq[T]) (T, bool) {
	var maxVal T
	var ok bool
	for t := range it {
		if !ok {
			maxVal = t
			ok = true
		} else {
			if t > maxVal {
				maxVal = t
			}
		}
	}
	return maxVal, ok
}

func Min[T cmp.Ordered](it iter.Seq[T]) (T, bool) {
	var minVal T
	var ok bool
	for t := range it {
		if !ok {
			minVal = t
			ok = true
		} else {
			if t < minVal {
				minVal = t
			}
		}
	}
	return minVal, ok
}

func MaxByKey[T any, R cmp.Ordered](it iter.Seq[T], keyFn func(T) R) (T, bool) {
	var maxVal T
	var maxKey R
	var ok bool
	var key R
	for t := range it {
		key = keyFn(t)
		if !ok {
			ok = true
			maxVal, maxKey = t, key
		} else {
			if key > maxKey {
				maxVal, maxKey = t, key
			}
		}
	}
	return maxVal, ok
}

func MaxBy[T any](it iter.Seq[T], compareFn func(T, T) int) (T, bool) {
	var maxVal T
	var ok bool
	for t := range it {
		if !ok {
			maxVal = t
			ok = true
		} else {
			if compareFn(t, maxVal) > 0 {
				maxVal = t
			}
		}
	}
	return maxVal, ok
}

func MinByKey[T any, R cmp.Ordered](it iter.Seq[T], keyFn func(T) R) (T, bool) {
	var minVal T
	var minKey R
	var ok bool
	var key R
	for t := range it {
		key = keyFn(t)
		if !ok {
			ok = true
			minVal, minKey = t, key
		} else {
			if key < minKey {
				minVal, minKey = t, key
			}
		}
	}
	return minVal, ok
}

func MinBy[T any](it iter.Seq[T], compareFn func(T, T) int) (T, bool) {
	var minVal T
	var ok bool
	for t := range it {
		if !ok {
			minVal = t
			ok = true
		} else {
			if compareFn(t, minVal) < 0 {
				minVal = t
			}
		}
	}
	return minVal, ok
}

func Rev[T any](it iter.Seq[T]) iter.Seq[T] {
	temp := slices.Collect(it)
	return func(yield func(T) bool) {
		for i := len(temp) - 1; i >= 0; i-- {
			if !yield(temp[i]) {
				return
			}
		}
	}
}

func Unzip[T, O any](it iter.Seq2[T, O]) (iter.Seq[T], iter.Seq[O]) {
	var it1 iter.Seq[T]
	var it2 iter.Seq[O]
	it1 = func(yield func(T) bool) {
		for t := range it {
			if !yield(t) {
				return
			}
		}
	}
	it2 = func(yield func(O) bool) {
		for _, o := range it {
			if !yield(o) {
				return
			}
		}
	}
	return it1, it2
}

func Cloned[T any](it iter.Seq[*T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range it {
			if !yield(*t) {
				return
			}
		}
	}
}

func Cycle[T any](it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			for t := range it {
				if !yield(t) {
					return
				}
			}
		}
	}
}

func Sum[T cmp.Ordered](it iter.Seq[T]) T {
	var sum T
	for t := range it {
		sum += t
	}
	return sum
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Product[T numeric](it iter.Seq[T]) T {
	var prod T = 1
	for t := range it {
		prod *= t
	}
	return prod
}

func Cmp[T cmp.Ordered](it, other iter.Seq[T]) int {
	next, stop := iter.Pull(other)
	defer stop()
	var ok bool
	var o T
	var comp int
	for t := range it {
		o, ok = next()
		if !ok {
			return 1
		}
		if comp = cmp.Compare(t, o); comp != 0 {
			return comp
		}
	}
	o, ok = next()
	if ok {
		return -1
	}
	return 0
}

func Ne[T cmp.Ordered](it, other iter.Seq[T]) bool {
	return Cmp(it, other) != 0
}

func Lt[T cmp.Ordered](it, other iter.Seq[T]) bool {
	return Cmp(it, other) < 0
}

func Le[T cmp.Ordered](it, other iter.Seq[T]) bool {
	return Cmp(it, other) <= 0
}

func Gt[T cmp.Ordered](it, other iter.Seq[T]) bool {
	return Cmp(it, other) > 0
}

func Ge[T cmp.Ordered](it, other iter.Seq[T]) bool {
	return Cmp(it, other) == 0
}

func IsSorted[T cmp.Ordered](it iter.Seq[T]) bool {
	index := uint(0)
	var prev T
	var comp, curComp int
	for t := range it {
		if index == 0 {
			prev = t
			index++
			continue
		}
		curComp, prev = cmp.Compare(prev, t), t
		if comp == 0 && curComp != 0 {
			comp = curComp
		}
		if comp*curComp < 0 {
			return false
		}
		index++
	}
	return true
}

func IsSortedBy[T any](it iter.Seq[T], compareFn func(T, T) int) bool {
	index := uint(0)
	var prev T
	var comp, curComp int
	for t := range it {
		if index == 0 {
			prev = t
			index++
			continue
		}
		curComp, prev = compareFn(prev, t), t
		if comp == 0 && curComp != 0 {
			comp = curComp
		}
		if comp*curComp < 0 {
			return false
		}
		index++
	}
	return true
}

func IsSortedByKey[T any, K cmp.Ordered](it iter.Seq[T], keyFn func(T) K) bool {
	index := uint(0)
	var prev K
	var comp, curComp int
	for t := range it {
		if index == 0 {
			prev = keyFn(t)
			index++
			continue
		}
		curComp, prev = cmp.Compare(prev, keyFn(t)), keyFn(t)
		if comp == 0 && curComp != 0 {
			comp = curComp
		}
		if comp*curComp < 0 {
			return false
		}
		index++
	}
	return true
}
