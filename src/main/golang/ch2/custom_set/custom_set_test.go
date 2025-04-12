package custom_set

import "testing"

func TestCustomSet(t *testing.T) {
	if s := New[string](); s == nil {
		t.Errorf("New[string]() returned nil %+v", s)
	}
}
