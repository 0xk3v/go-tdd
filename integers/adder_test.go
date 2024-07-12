package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(1, 3)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}
