package say

import "testing"

var extension = new(SAY)

func TestHello(t *testing.T) {

	got := extension.Hello("Daniel")
	want := "Hello Daniel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
