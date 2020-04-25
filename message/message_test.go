package message

import "testing"

func TestGreet(t *testing.T) {
	actual := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if actual != expect {
		t.Errorf("Got: %q, wanted %q", actual, expect)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "Test", expect: "Hello, Test!\n"},
		{input: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		actual := Greet(s.input)
		if actual != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected: %q, actual: %q",
				s.input, s.expect, actual)
		}
	}
}

func TestDepart(t *testing.T) {
	actual := depart("Gopher")
	expect := "Goodbye, Gopher\n"
	if actual != expect {
		t.Errorf("Got: %q, wanted %q", actual, expect)
	}
}
