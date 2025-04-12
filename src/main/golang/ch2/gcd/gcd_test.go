package gcd

import "testing"

func TestGCD(t *testing.T) {
	for _, tc := range testcases {
		got := GCD(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("GCD(%v, %v) = %d, want %d", tc.a, tc.b, got, tc.want)
		} else {
			t.Logf("passed GCD(%v, %v) = %d", tc.a, tc.b, got)
		}
	}
}

var testcases = []struct {
	a, b, want uint64
}{
	{0, 0, 0},
	{0, 7, 7},
	{7, 0, 7},
	{42, 36, 6},
	{24, 48, 24},
	{31, 31, 31},
	{18446744073709551615, 18446744073709551612, 3},
	{10001, 23435, 1},
}
