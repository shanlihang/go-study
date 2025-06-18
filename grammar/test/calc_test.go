package test

import (
	"testing"
)
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	t.Log()
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}
