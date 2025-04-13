package custom_iter

import (
	"fmt"
	"iter"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tcs := []struct {
		input []int
		want  uint
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
		{[]int{}, 0},
		{nil, 0},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("count %v", tc.input), func(t *testing.T) {
			got := Count(slices.Values(tc.input))
			if got != tc.want {
				t.Errorf("got: %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	tcs := []struct {
		input []int
		want1 int
		want2 bool
	}{
		{[]int{1, 3, 2, 4}, 4, true},
		{[]int{}, 0, false},
		{nil, 0, false},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("last %v", tc.input), func(t *testing.T) {
			got1, got2 := Last(slices.Values(tc.input))
			if got1 != tc.want1 || got2 != tc.want2 {
				t.Errorf("got: %v %v want %v %v\n", got1, got2, tc.want1, tc.want2)
			}
		})
	}
}

func TestNth(t *testing.T) {
	tcs := []struct {
		input []int
		n     uint
		want1 int
		want2 bool
	}{
		{[]int{1, 3, 2, 4}, 0, 1, true},
		{[]int{1, 3, 2, 4}, 1, 3, true},
		{[]int{1, 3, 2, 4}, 4, 0, false},
		{nil, 0, 0, false},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("nth %v %v", tc.input, tc.n), func(t *testing.T) {
			got1, got2 := Nth(slices.Values(tc.input), tc.n)
			if got1 != tc.want1 || got2 != tc.want2 {
				t.Errorf("got: %v %v want %v %v\n", got1, got2, tc.want1, tc.want2)
			}
		})
	}
}

func TestStepBy(t *testing.T) {
	tcs := []struct {
		input []int
		step  uint
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, []int{1, 3, 5, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("stepby %v %v", tc.input, tc.step), func(t *testing.T) {
			got := slices.Collect(StepBy(slices.Values(tc.input), tc.step))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v\n", got, tc.want)
			}
		})
	}
}

func TestChain(t *testing.T) {
	tcs := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{nil, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{1}, []int{1}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("chain %v %v", tc.input1, tc.input2), func(t *testing.T) {
			got := slices.Collect(Chain(slices.Values(tc.input1), slices.Values(tc.input2)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	tcs := []struct {
		input1 []int
		input2 []rune
		want1  []int
		want2  []rune
	}{
		{[]int{1, 2, 3, 4, 5}, []rune{'a', 'b', 'c'}, []int{1, 2, 3}, []rune{'a', 'b', 'c'}},
		{[]int{1, 2, 3}, []rune{'a', 'b', 'c', 'd'}, []int{1, 2, 3}, []rune{'a', 'b', 'c'}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("zip %v %v", tc.input1, tc.input2), func(t *testing.T) {
			var got1 []int
			var got2 []rune
			for x, y := range Zip(slices.Values(tc.input1), slices.Values(tc.input2)) {
				got1 = append(got1, x)
				got2 = append(got2, y)
			}
			got1 = got1[:len(got1):len(got1)]
			got2 = got2[:len(got2):len(got2)]
			if !reflect.DeepEqual(got1, tc.want1) || !reflect.DeepEqual(got2, tc.want2) {
				t.Errorf("got: %v %v want: %v %v\n", got1, got2, tc.want1, tc.want2)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tcs := []struct {
		input []int
		want  string
	}{
		{[]int{1, 2, 3}, "123"},
		{[]int{}, ""},
	}
	var builder strings.Builder
	do := func(i int) {
		builder.WriteString(strconv.Itoa(i))
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("forEach %v", tc.input), func(t *testing.T) {
			ForEach(slices.Values(tc.input), do)
			got := builder.String()
			builder.Reset()
			if got != tc.want {
				t.Errorf("got %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tcs := []struct {
		input []int
		pred  func(int) bool
		want  bool
	}{
		{[]int{}, func(i int) bool { return true }, true},
		{[]int{1}, func(i int) bool { return true }, true},
		{[]int{1, 2}, func(i int) bool { return i%2 == 0 }, false},
		{[]int{1, 2, 3}, func(i int) bool { return i < 10 }, true},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("all %v", tc.input), func(t *testing.T) {
			got := All(slices.Values(tc.input), tc.pred)
			if got != tc.want {
				t.Errorf("got %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	tcs := []struct {
		input []int
		pred  func(int) bool
		want  bool
	}{
		{[]int{}, func(i int) bool { return true }, false},
		{[]int{1}, func(i int) bool { return true }, true},
		{[]int{1, 2}, func(i int) bool { return i%2 == 0 }, true},
		{[]int{1, 2, 3}, func(i int) bool { return i > 10 }, false},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("all %v", tc.input), func(t *testing.T) {
			got := Any(slices.Values(tc.input), tc.pred)
			if got != tc.want {
				t.Errorf("got %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestByRef(t *testing.T) {
	type pair struct{ x, y int }
	tcs := []struct {
		input []pair
		want  []*pair
	}{
		{[]pair{{3, 5}, {4, 2}}, []*pair{{3, 5}, {4, 2}}},
		{nil, nil},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("by ref %v", tc.input), func(t *testing.T) {
			got := slices.Collect(ByRef(slices.Values(tc.input)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestCollectIntoSlice(t *testing.T) {
	tcs := []struct {
		input iter.Seq[int]
		want  []int
	}{
		{func(yield func(int) bool) {
			for i := range 5 {
				if !yield((i + 1) * 2) {
					return
				}
			}
		}, []int{2, 4, 6, 8, 10}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("collect into slice %v", tc.want), func(t *testing.T) {
			got := CollectIntoSlice(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v want %v\n", got, tc.want)
			}
		})
	}
}

func TestPartitionIntoSlices(t *testing.T) {
	tcs := []struct {
		input        iter.Seq[int]
		pred         func(int) bool
		want1, want2 []int
	}{
		{func(yield func(int) bool) {
			for i := range 10 {
				if !yield(i + 1) {
					return
				}
			}
		}, func(i int) bool { return i%2 == 0 }, []int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}},
		{func(yield func(int) bool) {
			for i := range 10 {
				if !yield(10 - i) {
					return
				}
			}
		}, func(i int) bool { return i > 7 }, []int{10, 9, 8}, []int{7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("partition into slices %v %v", tc.want1, tc.want2), func(t *testing.T) {
			got1, got2 := PartitionIntoSlices(tc.input, tc.pred)
			if !reflect.DeepEqual(got1, tc.want1) || !reflect.DeepEqual(got2, tc.want2) {
				t.Errorf("got %v %v want %v %v\n", got1, got2, tc.want1, tc.want2)
			}
		})
	}
}

// TODO: Add tests for TryFold and the rest
