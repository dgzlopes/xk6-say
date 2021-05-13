package say

import "go.k6.io/k6/js/modules"

func init() {
	modules.Register("k6/x/say", new(SAY))
}

// SAY is the k6 say extension.
type SAY struct{}

// Hello says hello to someone
func (*SAY) Hello(name string) string {
	return "Hello " + name
}
