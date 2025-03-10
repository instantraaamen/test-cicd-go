package operations

import "testing"

func TestAdd(t *testing.T) {
	got := Add(10, 5)
	want := 15
	if got != want {
		t.Errorf("Add(10,5) = %d; want %d", got, want)
	}
}
