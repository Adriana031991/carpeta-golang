package main

import "testing"

func TestMiSumauno(t *testing.T) {
	v := miSuma(1, 3)
	if v != 4 {
		t.Error("Expected", 4, "Got", miSuma(1, 3))
	}
}
