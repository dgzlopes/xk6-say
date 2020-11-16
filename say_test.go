package say

import "testing"

var plugin = new(SAY)

func TestHello(t *testing.T) {

	got := plugin.Hello("Daniel")
	want := "Hello Daniel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
