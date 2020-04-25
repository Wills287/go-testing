package message

import "testing"

func TestGreet(t *testing.T) {
	actual := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if actual != expect {
		t.Errorf("Got: %q, wanted %q", actual, expect)
	}
}

func TestDepart(t *testing.T) {
	actual := depart("Gopher")
	expect := "Goodbye, Gopher\n"
	if actual != expect {
		t.Errorf("Got: %q, wanted %q", actual, expect)
	}
}
