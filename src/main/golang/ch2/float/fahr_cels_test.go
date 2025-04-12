package float

import "testing"

type testCase[E Temperature] struct {
	description string
	input       E
	output      E
}

func TestFahrToCels(t *testing.T) {
	for _, tc1 := range f2c1 {
		t.Run(tc1.description, func(t *testing.T) {
			got, err := FahrToCels(tc1.input)
			if got != tc1.output || err != nil {
				t.Errorf("%s got %v, error %+v, want %v", tc1.description, got, err, tc1.output)
			} else {
				t.Logf("passed %s", tc1.description)
			}
		})
	}
	for _, tc2 := range f2c2 {
		t.Run(tc2.description, func(t *testing.T) {
			got, err := FahrToCels(tc2.input)
			if got != tc2.output || err != nil {
				t.Errorf("%s got %v, err %+v, want %v", tc2.description, got, err, tc2.output)
			} else {
				t.Logf("passed %s", tc2.description)
			}
		})
	}
	for _, tc3 := range f2c3 {
		t.Run(tc3.description, func(t *testing.T) {
			got, err := FahrToCels(tc3.input)
			if got != tc3.output || err != nil {
				t.Errorf("%s got %v, error %+v, want %v", tc3.description, got, err, tc3.output)
			} else {
				t.Logf("passed %s", tc3.description)
			}
		})
	}
}

var f2c1 = []testCase[float64]{
	{
		description: "25.7F->-3.5",
		input:       25.7,
		output:      -3.5,
	},
}

var f2c2 = []testCase[int]{
	{
		description: "23F->-5C",
		input:       23,
		output:      -5,
	},
}

var f2c3 = []testCase[uint]{
	{
		description: "50F->10C",
		input:       50,
		output:      10,
	},
}

func TestCelsToFahr(t *testing.T) {
	for _, tc1 := range c2f1 {
		t.Run(tc1.description, func(t *testing.T) {
			got, err := CelsToFahr(tc1.input)
			if got != tc1.output || err != nil {
				t.Errorf("%s got %v, error %+v, want %v", tc1.description, got, err, tc1.output)
			} else {
				t.Logf("passed %s", tc1.description)
			}
		})
	}
	for _, tc2 := range c2f2 {
		t.Run(tc2.description, func(t *testing.T) {
			got, err := CelsToFahr(tc2.input)
			if got != tc2.output || err != nil {
				t.Errorf("%s got %v, error %+v, want %v", tc2.description, got, err, tc2.output)
			} else {
				t.Logf("passed %s", tc2.description)
			}
		})
	}
	for _, tc3 := range c2f3 {
		t.Run(tc3.description, func(t *testing.T) {
			got, err := CelsToFahr(tc3.input)
			if got != tc3.output || err != nil {
				t.Errorf("%s got %v, error %+v, want %v", tc3.description, got, err, tc3.output)
			} else {
				t.Logf("passed %s", tc3.description)
			}
		})
	}
}

var c2f1 = []testCase[float64]{
	{
		description: "25.3C->77.54F",
		input:       25.3,
		output:      77.54,
	},
}

var c2f2 = []testCase[int]{
	{
		description: "-10C->14F",
		input:       -10,
		output:      14,
	},
}

var c2f3 = []testCase[uint]{
	{
		description: "0C->32F",
		input:       0,
		output:      32,
	},
}
