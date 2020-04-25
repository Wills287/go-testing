package main_test

import "testing"

func TestAddition(t *testing.T) {
	actual := 2 + 2
	expected := 4
	if actual != expected {
		t.Errorf("Got: '%v', wanted: '%v'", actual, expected)
	}
}

func TestSubtraction(t *testing.T) {
	actual := 10 - 5
	expected := 4
	if actual != expected {
		t.Errorf("Got: '%v', wanted: '%v'", actual, expected)
	}
}
